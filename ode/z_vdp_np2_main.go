// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"testing"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
	"github.com/cpmech/gosl/ode"
)

func main() {

	// start mpi
	mpi.Start()
	defer mpi.Stop()

	//check number of processors
	if mpi.WorldRank() == 0 {
		chk.Verbose = true
		chk.PrintTitle("Hairer-Wanner VII-p5 Eq.(1.5) Van der Pol's Equation (Distr=true)")
	}
	if mpi.WorldSize() != 2 {
		if mpi.WorldRank() == 0 {
			io.Pf("ERROR: this test needs 2 processors (run with mpi -np 2)\n")
		}
		return
	}

	// communicator
	comm := mpi.NewCommunicator(nil)

	// dy/dx function
	eps := 1.0e-6
	w := la.NewVector(2) // workspace
	fcn := func(f la.Vector, dx, x float64, y la.Vector) {
		w.Fill(0)
		switch comm.Rank() {
		case 0:
			w[0] = y[1]
		case 1:
			w[1] = ((1.0-y[0]*y[0])*y[1] - y[0]) / eps
		}
		comm.AllReduceSum(f, w)
	}

	// Jacobian
	jac := func(dfdy *la.Triplet, dx, x float64, y la.Vector) {
		if dfdy.Max() == 0 {
			dfdy.Init(2, 2, 4)
		}
		dfdy.Start()
		switch comm.Rank() {
		case 0:
			dfdy.Put(0, 0, 0.0)
			dfdy.Put(0, 1, 1.0)
		case 1:
			dfdy.Put(1, 0, (-2.0*y[0]*y[1]-1.0)/eps)
			dfdy.Put(1, 1, (1.0-y[0]*y[0])/eps)
		}
	}

	// initial values
	xb := 2.0
	ndim := 2
	y := la.NewVectorSlice([]float64{2.0, -0.6})

	// configurations
	conf := ode.NewConfig("radau5", "mumps", comm)
	conf.SetStepOut(true, nil)
	conf.SetTol(1e-4)

	// solver
	sol := ode.NewSolver(ndim, conf, fcn, jac, nil)

	// solve
	sol.Solve(y, 0, xb)

	// only root
	if mpi.WorldRank() == 0 {

		//check
		tst := new(testing.T)
		chk.Int(tst, "number of F evaluations ", sol.Stat.Nfeval, 2233)
		chk.Int(tst, "number of J evaluations ", sol.Stat.Njeval, 160)
		chk.Int(tst, "total number of steps   ", sol.Stat.Nsteps, 280)
		chk.Int(tst, "number of accepted steps", sol.Stat.Naccepted, 241)
		chk.Int(tst, "number of rejected steps", sol.Stat.Nrejected, 7)
		chk.Int(tst, "number of decompositions", sol.Stat.Ndecomp, 251)
		chk.Int(tst, "number of lin solutions ", sol.Stat.Nlinsol, 663)
		chk.Int(tst, "max number of iterations", sol.Stat.Nitmax, 6)
		chk.String(tst, sol.Stat.LsKind, "mumps")
	}
}
