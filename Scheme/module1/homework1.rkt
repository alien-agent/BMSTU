#lang racket
(require racket/date)

; task 1
(define (my-odd? x) (odd? x))
(define (my-even? x) (even? x))

(define (day-of-week day month year)
  (define local-secs (find-seconds 0 0 0 day month year #t))
  (define my-date (seconds->date local-secs))
  (date-week-day my-date))

; task 2
(define (solve a b c)
  (define disc (- (* b b) (* 4.0 a c)))
  (cond
    [(< disc 0) (list)]
    [(= disc 0) (list (/ (- b) (* 2.0 a)))]
    [(> disc 0) (list (/ (+ (- b) (sqrt disc)) (* 2.0 a)) (/ (- (- b) (sqrt disc)) (* 2.0 a)))]))

; task 3.1
(define (gcd a b)
  (cond
    [(> a b) (gcd b (- a b))]
    [(< a b) (gcd a (- b a))]
    [else a]))

; task 3.2
(define (lcm a b) (/ (* a b) (gcd a b)))

; task 3.3
(define (helper x k)
  (if (= x k)
      #t
      (if (= (remainder x k) 0) #f
          (helper x (+ k 1)))))

(define (prime? x )
  (cond
    (( = x 1 ) #t)
    (( = x 2 ) #t)
    ( else (helper x 2 ) )))