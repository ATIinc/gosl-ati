# Gosl. num/qpck. Wrapper to QUADPACK

[![PkgGoDev](https://pkg.go.dev/badge/github.com/ATIinc/gosl-ati/num/qpck)](https://pkg.go.dev/github.com/ATIinc/gosl-ati/num/qpck)

This package wraps the QUADPACK package written in FORTRAN to compute definite one-dimensional
integrals.

The code here was obtained from the [SciPy version](https://github.com/scipy/scipy/tree/master/scipy/integrate/quadpack).

## Original Readme

```
QUADPACK is a FORTRAN subroutine package for the numerical
computation of definite one-dimensional integrals. It originated
from a joint project of R. Piessens and E. de Doncker (Appl.
Math. and Progr. Div.- K.U.Leuven, Belgium), C. Ueberhuber (Inst.
Fuer Math.- Techn.U.Wien, Austria), and D. Kahaner (Nation. Bur.
of Standards- Washington D.C., U.S.A.).
The routine names for the DOUBLE PRECISION versions are preceded
by the letter D.

- QNG  : Is a simple non-adaptive automatic integrator, based on
         a sequence of rules with increasing degree of algebraic
         precision (Patterson, 1968).

- QAG  : Is a simple globally adaptive integrator using the
         strategy of Aind (Piessens, 1973). It is possible to
         choose between 6 pairs of Gauss-Kronrod quadrature
         formulae for the rule evaluation component. The pairs
         of high degree of precision are suitable for handling
         integration difficulties due to a strongly oscillating
         integrand.

- QAGS : Is an integrator based on globally adaptive interval
         subdivision in connection with extrapolation (de Doncker,
         1978) by the Epsilon algorithm (Wynn, 1956).

- QAGP : Serves the same purposes as QAGS, but also allows
         for eventual user-supplied information, i.e. the
         abscissae of internal singularities, discontinuities
         and other difficulties of the integrand function.
         The algorithm is a modification of that in QAGS.

- QAGI : Handles integration over infinite intervals. The
         infinite range is mapped onto a finite interval and
         then the same strategy as in QAGS is applied.

- QAWO : Is a routine for the integration of COS(OMEGA*X)*F(X)
         or SIN(OMEGA*X)*F(X) over a finite interval (A,B).
         OMEGA is is specified by the user
         The rule evaluation component is based on the
         modified Clenshaw-Curtis technique.
         An adaptive subdivision scheme is used connected with
         an extrapolation procedure, which is a modification
         of that in QAGS and provides the possibility to deal
         even with singularities in F.

- QAWF : Calculates the Fourier cosine or Fourier sine
         transform of F(X), for user-supplied interval (A,
         INFINITY), OMEGA, and F. The procedure of QAWO is
         used on successive finite intervals, and convergence
         acceleration by means of the Epsilon algorithm (Wynn,
         1956) is applied to the series of the integral
         contributions.

- QAWS : Integrates W(X)*F(X) over (A,B) with A.LT.B finite,
         and   W(X) = ((X-A)**ALFA)*((B-X)**BETA)*V(X)
         where V(X) = 1 or LOG(X-A) or LOG(B-X)
                        or LOG(X-A)*LOG(B-X)
         and   ALFA.GT.(-1), BETA.GT.(-1).
         The user specifies A, B, ALFA, BETA and the type of
         the function V.
         A globally adaptive subdivision strategy is applied,
         with modified Clenshaw-Curtis integration on the
         subintervals which contain A or B.

- QAWC : Computes the Cauchy Principal Value of F(X)/(X-C)
         over a finite interval (A,B) and for
         user-determined C.
         The strategy is globally adaptive, and modified
         Clenshaw-Curtis integration is used on the subranges
         which contain the point X = C.

   Each of the routines above also has a "more detailed" version
with a name ending in E, as QAGE.  These provide more
information and control than the easier versions.


   The preceeding routines are all automatic.  That is, the user
inputs his problem and an error tolerance.  The routine
attempts to perform the integration to within the requested
absolute or relative error.
   There are, in addition, a number of non-automatic integrators.
These are most useful when the problem is such that the
user knows that a fixed rule will provide the accuracy
required.  Typically they return an error estimate but make
no attempt to satisfy any particular input error request.

  QK15 QK21 QK31 QK41 QK51 QK61
       Estimate the integral on [a,b] using 15, 21,..., 61
       point rule and return an error estimate.
  QK15I 15 point rule for (semi)infinite interval.
  QK15W 15 point rule for special singular weight functions.
  QC25C 25 point rule for Cauchy Principal Values
  QC25F 25 point rule for sin/cos integrand.
  QMOMO Integrates k-th degree Chebychev polynomial times
        function with various explicit singularities.

Support functions from linpack, slatec, and blas have been omitted
by default but can be obtained by asking.  For example, suppose you
already have installed linpack and the blas, but not slatec.  Then
use a request like  "send dqag from quadpack slatec".

[see also toms/691]
```

## Original doc file

This version obtained on 1 Jun 84 from (kahaner@nbs-sdc)
David K. Kahaner
Scientific Computing Division
National Bureau of Standards
Washington DC 20234

WARNING: the calling sequences here differ from those in the book:
R. Piessens, E. deDoncker-Kapenga, C. Uberhuber, D. Kahaner
Quadpack: a Subroutine Package for Automatic Integration
Springer Verlag, 1983. Series in Computational Mathematics v.1

```
See also:
  R. Piessens, E. deDoncker-Kapenga, C. Uberhuber, D. Kahaner
  Quadpack: a Subroutine Package for Automatic Integration
  Springer Verlag, 1983.    Series in Computational Mathematics v.1
  515.43/Q1S  100394Z

      SUBROUTINE QPDOC
C***BEGIN PROLOGUE  QPDOC
C***DATE WRITTEN   810401   (YYMMDD)
C***REVISION DATE  840417   (YYMMDD)
C***CATEGORY NO.  H2
C***KEYWORDS  SURVEY OF INTEGRATORS, GUIDELINES FOR SELECTION,QUADPACK
C***AUTHOR  PIESSENS, ROBERT(APPL. MATH. AND PROGR. DIV.- K.U.LEUVEN)
C           DE DONKER, ELISE(APPL. MATH. AND PROGR. DIV.- K.U.LEUVEN
C           KAHANER,DAVID(NATIONAL BUREAU OF STANDARDS)
C***PURPOSE  QUADPACK documentation routine.
C***DESCRIPTION
C 1. Introduction
C    ------------
C    QUADPACK is a FORTRAN subroutine package for the numerical
C    computation of definite one-dimensional integrals. It originated
C    from a joint project of R. Piessens and E. de Doncker (Appl.
C    Math. and Progr. Div.- K.U.Leuven, Belgium), C. Ueberhuber (Inst.
C    Fuer Math.- Techn.U.Wien, Austria), and D. Kahaner (Nation. Bur.
C    of Standards- Washington D.C., U.S.A.).
C    Documentation routine QPDOC describes the package in the form it
C    was released from A.M.P.D.- Leuven, for adherence to the SLATEC
C    library in May 1981. Apart from a survey of the integrators, some
C    guidelines will be given in order to help the QUADPACK user with
C    selecting an appropriate routine or a combination of several
C    routines for handling his problem.
C
C    In the LONG DESCRIPTION of QPDOC it is demonstrated how to call
C    the integrators, by means of small example calling programs.
C
C    For precise guidelines involving the use of each routine in
C    particular, we refer to the extensive introductory comments
C    within each routine.
C
C 2. Survey
C    ------
C    The following list gives an overview of the QUADPACK integrators.
C    The routine names for the DOUBLE PRECISION versions are preceded
C    by the letter D.
C
C    - QNG  : Is a simple non-adaptive automatic integrator, based on
C             a sequence of rules with increasing degree of algebraic
C             precision (Patterson, 1968).
C
C    - QAG  : Is a simple globally adaptive integrator using the
C             strategy of Aind (Piessens, 1973). It is possible to
C             choose between 6 pairs of Gauss-Kronrod quadrature
C             formulae for the rule evaluation component. The pairs
C             of high degree of precision are suitable for handling
C             integration difficulties due to a strongly oscillating
C             integrand.
C
C    - QAGS : Is an integrator based on globally adaptive interval
C             subdivision in connection with extrapolation (de Doncker,
C             1978) by the Epsilon algorithm (Wynn, 1956).
C
C    - QAGP : Serves the same purposes as QAGS, but also allows
C             for eventual user-supplied information, i.e. the
C             abscissae of internal singularities, discontinuities
C             and other difficulties of the integrand function.
C             The algorithm is a modification of that in QAGS.
C
C    - QAGI : Handles integration over infinite intervals. The
C             infinite range is mapped onto a finite interval and
C             then the same strategy as in QAGS is applied.
C
C    - QAWO : Is a routine for the integration of COS(OMEGA*X)*F(X)
C             or SIN(OMEGA*X)*F(X) over a finite interval (A,B).
C             OMEGA is is specified by the user
C             The rule evaluation component is based on the
C             modified Clenshaw-Curtis technique.
C             An adaptive subdivision scheme is used connected with
C             an extrapolation procedure, which is a modification
C             of that in QAGS and provides the possibility to deal
C             even with singularities in F.
C
C    - QAWF : Calculates the Fourier cosine or Fourier sine
C             transform of F(X), for user-supplied interval (A,
C             INFINITY), OMEGA, and F. The procedure of QAWO is
C             used on successive finite intervals, and convergence
C             acceleration by means of the Epsilon algorithm (Wynn,
C             1956) is applied to the series of the integral
C             contributions.
C
C    - QAWS : Integrates W(X)*F(X) over (A,B) with A.LT.B finite,
C             and   W(X) = ((X-A)**ALFA)*((B-X)**BETA)*V(X)
C             where V(X) = 1 or LOG(X-A) or LOG(B-X)
C                            or LOG(X-A)*LOG(B-X)
C             and   ALFA.GT.(-1), BETA.GT.(-1).
C             The user specifies A, B, ALFA, BETA and the type of
C             the function V.
C             A globally adaptive subdivision strategy is applied,
C             with modified Clenshaw-Curtis integration on the
C             subintervals which contain A or B.
C
C    - QAWC : Computes the Cauchy Principal Value of F(X)/(X-C)
C             over a finite interval (A,B) and for
C             user-determined C.
C             The strategy is globally adaptive, and modified
C             Clenshaw-Curtis integration is used on the subranges
C             which contain the point X = C.
C
C  Each of the routines above also has a "more detailed" version
C    with a name ending in E, as QAGE.  These provide more
C    information and control than the easier versions.
C
C
C   The preceeding routines are all automatic.  That is, the user
C      inputs his problem and an error tolerance.  The routine
C      attempts to perform the integration to within the requested
C      absolute or relative error.
C   There are, in addition, a number of non-automatic integrators.
C      These are most useful when the problem is such that the
C      user knows that a fixed rule will provide the accuracy
C      required.  Typically they return an error estimate but make
C      no attempt to satisfy any particular input error request.
C
C      QK15
C      QK21
C      QK31
C      QK41
C      QK51
C      QK61
C           Estimate the integral on [a,b] using 15, 21,..., 61
C           point rule and return an error estimate.
C      QK15I 15 point rule for (semi)infinite interval.
C      QK15W 15 point rule for special singular weight functions.
C      QC25C 25 point rule for Cauchy Principal Values
C      QC25F 25 point rule for sin/cos integrand.
C      QMOMO Integrates k-th degree Chebychev polynomial times
C            function with various explicit singularities.
C
C 3. Guidelines for the use of QUADPACK
C    ----------------------------------
C    Here it is not our purpose to investigate the question when
C    automatic quadrature should be used. We shall rather attempt
C    to help the user who already made the decision to use QUADPACK,
C    with selecting an appropriate routine or a combination of
C    several routines for handling his problem.
C
C    For both quadrature over finite and over infinite intervals,
C    one of the first questions to be answered by the user is
C    related to the amount of computer time he wants to spend,
C    versus his -own- time which would be needed, for example, for
C    manual subdivision of the interval or other analytic
C    manipulations.
C
C    (1) The user may not care about computer time, or not be
C        willing to do any analysis of the problem. especially when
C        only one or a few integrals must be calculated, this attitude
C        can be perfectly reasonable. In this case it is clear that
C        either the most sophisticated of the routines for finite
C        intervals, QAGS, must be used, or its analogue for infinite
C        intervals, GAGI. These routines are able to cope with
C        rather difficult, even with improper integrals.
C        This way of proceeding may be expensive. But the integrator
C        is supposed to give you an answer in return, with additional
C        information in the case of a failure, through its error
C        estimate and flag. Yet it must be stressed that the programs
C        cannot be totally reliable.
C        ------
C
C    (2) The user may want to examine the integrand function.
C        If bad local difficulties occur, such as a discontinuity, a
C        singularity, derivative singularity or high peak at one or
C        more points within the interval, the first advice is to
C        split up the interval at these points. The integrand must
C        then be examinated over each of the subintervals separately,
C        so that a suitable integrator can be selected for each of
C        them. If this yields problems involving relative accuracies
C        to be imposed on -finite- subintervals, one can make use of
C        QAGP, which must be provided with the positions of the local
C        difficulties. However, if strong singularities are present
C        and a high accuracy is requested, application of QAGS on the
C        subintervals may yield a better result.
C
C        For quadrature over finite intervals we thus dispose of QAGS
C        and
C        - QNG for well-behaved integrands,
C        - QAG for functions with an oscillating behaviour of a non
C          specific type,
C        - QAWO for functions, eventually singular, containing a
C          factor COS(OMEGA*X) or SIN(OMEGA*X) where OMEGA is known,
C        - QAWS for integrands with Algebraico-Logarithmic end point
C          singularities of known type,
C        - QAWC for Cauchy Principal Values.
C
C        Remark
C        ------
C        On return, the work arrays in the argument lists of the
C        adaptive integrators contain information about the interval
C        subdivision process and hence about the integrand behaviour:
C        the end points of the subintervals, the local integral
C        contributions and error estimates, and eventually other
C        characteristics. For this reason, and because of its simple
C        globally adaptive nature, the routine QAG in particular is
C        well-suited for integrand examination. Difficult spots can
C        be located by investigating the error estimates on the
C        subintervals.
C
C        For infinite intervals we provide only one general-purpose
C        routine, QAGI. It is based on the QAGS algorithm applied
C        after a transformation of the original interval into (0,1).
C        Yet it may eventuate that another type of transformation is
C        more appropriate, or one might prefer to break up the
C        original interval and use QAGI only on the infinite part
C        and so on. These kinds of actions suggest a combined use of
C        different QUADPACK integrators. Note that, when the only
C        difficulty is an integrand singularity at the finite
C        integration limit, it will in general not be necessary to
C        break up the interval, as QAGI deals with several types of
C        singularity at the boundary point of the integration range.
C        It also handles slowly convergent improper integrals, on
C        the condition that the integrand does not oscillate over
C        the entire infinite interval. If it does we would advise
C        to sum succeeding positive and negative contributions to
C        the integral -e.g. integrate between the zeros- with one
C        or more of the finite-range integrators, and apply
C        convergence acceleration eventually by means of QUADPACK
C        subroutine QELG which implements the Epsilon algorithm.
C        Such quadrature problems include the Fourier transform as
C        a special case. Yet for the latter we have an automatic
C        integrator available, QAWF.
C
C***LONG DESCRIPTION
C 4. Example Programs
C    ----------------
C 4.1. Calling Program for QNG
C      -----------------------
C
C            REAL A,ABSERR,B,F,EPSABS,EPSREL,RESULT
C            INTEGER IER,NEVAL
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            CALL QNG(F,A,B,EPSABS,EPSREL,RESULT,ABSERR,NEVAL,IER)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = EXP(X)/(X*X+0.1E+01)
C            RETURN
C            END
C
C 4.2. Calling Program for QAG
C      -----------------------
C
C            REAL A,ABSERR,B,EPSABS,EPSREL,F,RESULT,WORK
C            INTEGER IER,IWORK,KEY,LAST,LENW,LIMIT,NEVAL
C            DIMENSION IWORK(100),WORK(400)
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            KEY = 6
C            LIMIT = 100
C            LENW = LIMIT*4
C            CALL QAG(F,A,B,EPSABS,EPSREL,KEY,RESULT,ABSERR,NEVAL,
C           *  IER,LIMIT,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 2.0E0/(2.0E0+SIN(31.41592653589793E0*X))
C            RETURN
C            END
C
C 4.3. Calling Program for QAGS
C      ------------------------
C
C            REAL A,ABSERR,B,EPSABS,EPSREL,F,RESULT,WORK
C            INTEGER IER,IWORK,LAST,LENW,LIMIT,NEVAL
C            DIMENSION IWORK(100),WORK(400)
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            LIMIT = 100
C            LENW = LIMIT*4
C            CALL QAGS(F,A,B,EPSABS,EPSREL,RESULT,ABSERR,NEVAL,IER,
C           *  LIMIT,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 0.0E0
C            IF(X.GT.0.0E0) F = 1.0E0/SQRT(X)
C            RETURN
C            END
C
C 4.4. Calling Program for QAGP
C      ------------------------
C
C            REAL A,ABSERR,B,EPSABS,EPSREL,F,POINTS,RESULT,WORK
C            INTEGER IER,IWORK,LAST,LENIW,LENW,LIMIT,NEVAL,NPTS2
C            DIMENSION IWORK(204),POINTS(4),WORK(404)
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            NPTS2 = 4
C            POINTS(1) = 1.0E0/7.0E0
C            POINTS(2) = 2.0E0/3.0E0
C            LIMIT = 100
C            LENIW = LIMIT*2+NPTS2
C            LENW = LIMIT*4+NPTS2
C            CALL QAGP(F,A,B,NPTS2,POINTS,EPSABS,EPSREL,RESULT,ABSERR,
C           *  NEVAL,IER,LENIW,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 0.0E+00
C            IF(X.NE.1.0E0/7.0E0.AND.X.NE.2.0E0/3.0E0) F =
C           *  ABS(X-1.0E0/7.0E0)**(-0.25E0)*
C           *  ABS(X-2.0E0/3.0E0)**(-0.55E0)
C            RETURN
C            END
C
C 4.5. Calling Program for QAGI
C      ------------------------
C
C            REAL ABSERR,BOUN,EPSABS,EPSREL,F,RESULT,WORK
C            INTEGER IER,INF,IWORK,LAST,LENW,LIMIT,NEVAL
C            DIMENSION IWORK(100),WORK(400)
C            EXTERNAL F
C            BOUN = 0.0E0
C            INF = 1
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            LIMIT = 100
C            LENW = LIMIT*4
C            CALL QAGI(F,BOUN,INF,EPSABS,EPSREL,RESULT,ABSERR,NEVAL,
C           *  IER,LIMIT,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 0.0E0
C            IF(X.GT.0.0E0) F = SQRT(X)*ALOG(X)/
C           *                   ((X+1.0E0)*(X+2.0E0))
C            RETURN
C            END
C
C 4.6. Calling Program for QAWO
C      ------------------------
C
C            REAL A,ABSERR,B,EPSABS,EPSREL,F,RESULT,OMEGA,WORK
C            INTEGER IER,INTEGR,IWORK,LAST,LENIW,LENW,LIMIT,MAXP1,NEVAL
C            DIMENSION IWORK(200),WORK(925)
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            OMEGA = 10.0E0
C            INTEGR = 1
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            LIMIT = 100
C            LENIW = LIMIT*2
C            MAXP1 = 21
C            LENW = LIMIT*4+MAXP1*25
C            CALL QAWO(F,A,B,OMEGA,INTEGR,EPSABS,EPSREL,RESULT,ABSERR,
C           *  NEVAL,IER,LENIW,MAXP1,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 0.0E0
C            IF(X.GT.0.0E0) F = EXP(-X)*ALOG(X)
C            RETURN
C            END
C
C 4.7. Calling Program for QAWF
C      ------------------------
C
C            REAL A,ABSERR,EPSABS,F,RESULT,OMEGA,WORK
C            INTEGER IER,INTEGR,IWORK,LAST,LENIW,LENW,LIMIT,LIMLST,
C           *  LST,MAXP1,NEVAL
C            DIMENSION IWORK(250),WORK(1025)
C            EXTERNAL F
C            A = 0.0E0
C            OMEGA = 8.0E0
C            INTEGR = 2
C            EPSABS = 1.0E-3
C            LIMLST = 50
C            LIMIT = 100
C            LENIW = LIMIT*2+LIMLST
C            MAXP1 = 21
C            LENW = LENIW*2+MAXP1*25
C            CALL QAWF(F,A,OMEGA,INTEGR,EPSABS,RESULT,ABSERR,NEVAL,
C           *  IER,LIMLST,LST,LENIW,MAXP1,LENW,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            IF(X.GT.0.0E0) F = SIN(50.0E0*X)/(X*SQRT(X))
C            RETURN
C            END
C
C 4.8. Calling Program for QAWS
C      ------------------------
C
C            REAL A,ABSERR,ALFA,B,BETA,EPSABS,EPSREL,F,RESULT,WORK
C            INTEGER IER,INTEGR,IWORK,LAST,LENW,LIMIT,NEVAL
C            DIMENSION IWORK(100),WORK(400)
C            EXTERNAL F
C            A = 0.0E0
C            B = 1.0E0
C            ALFA = -0.5E0
C            BETA = -0.5E0
C            INTEGR = 1
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            LIMIT = 100
C            LENW = LIMIT*4
C            CALL QAWS(F,A,B,ALFA,BETA,INTEGR,EPSABS,EPSREL,RESULT,
C           *  ABSERR,NEVAL,IER,LIMIT,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = SIN(10.0E0*X)
C            RETURN
C            END
C
C 4.9. Calling Program for QAWC
C      ------------------------
C
C            REAL A,ABSERR,B,C,EPSABS,EPSREL,F,RESULT,WORK
C            INTEGER IER,IWORK,LAST,LENW,LIMIT,NEVAL
C            DIMENSION IWORK(100),WORK(400)
C            EXTERNAL F
C            A = -1.0E0
C            B = 1.0E0
C            C = 0.5E0
C            EPSABS = 0.0E0
C            EPSREL = 1.0E-3
C            LIMIT = 100
C            LENW = LIMIT*4
C            CALL QAWC(F,A,B,C,EPSABS,EPSREL,RESULT,ABSERR,NEVAL,
C           *  IER,LIMIT,LENW,LAST,IWORK,WORK)
C      C  INCLUDE WRITE STATEMENTS
C            STOP
C            END
C      C
C            REAL FUNCTION F(X)
C            REAL X
C            F = 1.0E0/(X*X+1.0E-4)
C            RETURN
C            END
C***REFERENCES  (NONE)
C***ROUTINES CALLED  (NONE)
C***END PROLOGUE  QPDOC
```

## API

[Please see the documentation here](https://pkg.go.dev/github.com/ATIinc/gosl-ati/num/qpck)
