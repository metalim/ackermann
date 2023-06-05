# Ackermann function

Ackermann function A(m,n) is defined as follows:

* A(0,n) = n + 1
* A(m,0) = A(m-1, 1)
* A(m,n) = A(m-1, A(m, n-1))

By definition the function is computable and finite, but the only results computable in this universe are:
* A(0,n) ... A(3,n)
* A(4,0), A(4,1)
* A(5,0)

That's it. The rest is still computable, but will take more time, than universe exists. For instance result of A(4,2) takes 65536 bits (binary digits), and the only math operation in Ackermann function is n+1. So the best case estimate you can imagine is a counter with 65536 bits, coutning +1 per operation, until all but one bit are 1: 0b111111111...11111101 with 65536 digits

As for the rest of results — A(4,3+), A(5,1+), A(6+,n) — they cannot even fit in this universe in their numerical form, for instance value of A(4,3) takes 2<sup>65536</sup> bits, which is significantly more than number of atoms in the observable universe (2<sup>260</sup>)
