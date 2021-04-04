; Task 1

(define memoized-factorial
  (let ((memo '()))
    (lambda (n)
      (let ((memozed (assq n memo)))
        (if memozed
            (cadr memozed)
            (let ((new-value (if (< n 2)
                                 1
                                 (* (memoized-factorial (- n 1)) n))))
              (set! memo
                    (cons (list n new-value) memo))
              new-value))))))

; Task 2

(define-syntax lazy-cons
  (syntax-rules()
    ((_ a b) (cons a (delay b)))))

(define (lazy-car p)
  (car p))
(define (lazy-cdr p)
  (force (cdr p)))

(define (lazy-head xs k)
  (if (= k 0)
      '()
      (cons (lazy-car xs) (lazy-head (lazy-cdr xs) (- k 1)))))

(define (lazy-ref xs k)
  (if (= k 0)
      (lazy-car xs)
      (lazy-ref (lazy-cdr xs) (- k 1))))

(define (naturals start)
  (lazy-cons start (naturals (+ 1 start))))

(define (lazy-factorial n)
  (define (factorials k1 k2 )
    (lazy-cons k2 (factorials (+ k1 1) (* k2 (+ k1 1)))))
  (lazy-ref (factorials 0 1) n))

; Task 3
  
(define (read-words)
  (reader (read) '()))

(define (reader xs result)
  (if (eof-object? xs)
      result
      (if (number? xs)
          (reader (read) (append result (list (number->string xs))))
          (reader (read) (append result (list (symbol->string xs)))))))

; Task 4

(define-syntax meval
  (syntax-rules ()
  ((_ list) (eval list (interaction-environment)))))

(define (set-cadr! l v)
  (set-car! (cdr l) v))

(define (s->s arg)
  (if (symbol? arg)
      (symbol->string arg)
      (string->symbol arg)))

(define (strs->sym . strings)
  (s->s (apply string-append strings)))

(define-syntax gen-getter
  (syntax-rules ()
    ((_ name field) (meval (list 'define (list (strs->sym (s->s 'name) "-" (s->s 'field)) 'struct) '(cadr (assoc 'field (cdr struct))))))))

(define-syntax gen-setter
  (syntax-rules ()
    ((_ name field) (meval (list 'define (list (strs->sym "set-" (s->s 'name) "-" (s->s 'field) "!") 'struct 'val) '(set-cadr! (assoc 'field (cdr struct)) val))))))

(define-syntax gen-getters
  (syntax-rules ()
    ((_ name (field)) (gen-getter name field))
    ((_ name (field fields ...)) (begin (gen-getter name field) (gen-getters name (fields ...))))))

(define-syntax gen-setters
  (syntax-rules ()
    ((_ name (field)) (gen-setter name field))
    ((_ name (field fields ...)) (begin (gen-setter name field) (gen-setters name (fields ...))))))

(define-syntax gen-pred
  (syntax-rules ()
    ((_ name) (meval (list 'define (list (strs->sym (s->s 'name) "?") 'struct) '(and (list? struct) (symbol? (car struct)) (equal? (car struct) 'name)))))))

(define-syntax gen-record
  (syntax-rules ()
    ((_ (field)) (list (list 'list ''field (strs->sym (s->s 'field) "_arg"))))
    ((_ (field fields ...)) (append (list (list 'list ''field (strs->sym (s->s 'field) "_arg"))) (gen-record (fields ...))))
    ((_ name (fields ...)) (append '('name) (gen-record (fields ...))))))

(define-syntax gen-make
  (syntax-rules ()
    ((_ name (fields ...)) (meval (list 'define (list (strs->sym "make-" (s->s 'name)) (strs->sym (s->s 'fields) "_arg") ...) (append '(list) (gen-record name (fields ...))))))))

(define-syntax define-struct
  (syntax-rules ()
    ((_ name (fields ...))
     (begin
       (gen-getters name (fields ...))
       (gen-setters name (fields ...))
       (gen-pred name)
       (gen-make name (fields ...))))))

;(define-struct pos (row col)) ; Объявление типа pos
;(define p (make-pos 1 2))     ; Создание значения типа pos
;
;(pos? p)    ⇒ #t
;
;(pos-row p) ⇒ 1
;(pos-col p) ⇒ 2
;
;(set-pos-row! p 3) ; Изменение значения в поле row
;(set-pos-col! p 4) ; Изменение значения в поле col
;
;(pos-row p) ⇒ 3
;(pos-col p) ⇒ 4