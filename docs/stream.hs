stream :: (b->c) -> (b->c->Bool) -> (b->c->b) -> (b->a->b) -> b -> [a] -> [c]
stream next safe prod cons z (x:xs)
  = if   safe z y
    then y : stream next safe prod cons (prod z y) (x:xs)
    else     stream next safe prod cons (cons z x) xs
      where y = next z

type LFT = (Integer, Integer, Integer, Integer)
extr :: LFT -> Integer -> Rational
extr (q,r,s,t) x = ((fromInteger q) * (fromInteger x) + (fromInteger r)) /
                   ((fromInteger s) * (fromInteger x) + (fromInteger t))

unit :: LFT
unit = (1,0,0,1)

comp :: LFT -> LFT -> LFT
comp (q,r,s,t) (u,v,w,x) = (q*u+r*w,q*v+r*x,s*u+t*w,s*v+t*x)

pi = stream next safe prod cons init lfts where
  init      = unit
  lfts      = [(k, 4*k+2, 0, 2*k+1) | k<-[1..]]
  next z    = floor (extr z 3)
  safe z n  = (n == floor (extr z 4))
  prod z n  = comp (10, -10*n, 0, 1) z
  cons z z' = comp z z'
