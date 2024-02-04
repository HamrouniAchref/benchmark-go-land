package main

import (
	"fmt"
	"math/big"
	crand "crypto/rand"
	"crypto/sha256"
	"time"
//	"os"
	curve "github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/xuri/excelize/v2"

	//"unsafe"
	//"github.com/DmitriyVTitov/size"
	eddsa "github.com/consensys/gnark-crypto/ecc/bls12-381/twistededwards/eddsa"
//	"github.com/RadicalApp/libsignal-protocol-go/util/bytehelper"
 "strconv"
)

var x1 = []int{1,2,3,4,8,16,32}
var x2 = []int{64,128,256,512,1024,2048	}
var y1  [7][8]time.Duration	 
var y2  [6][8]time.Duration

func main() {
	_,_,G1,G2:=curve.Generators()
	for j, n := range x1 {
		fmt.Println("n=",n)
	benchmark_ME(n,G1,G2,j,1)
	benchmark_IsIn(n,G1,G2,j,1)
	benchmark_Pair(n,G1,G2,j,1)

	}
	fmt.Println("y1")
	fmt.Println(y1)


	for j, n := range x2 {
		fmt.Println("n=",n)
		benchmark_ME(n,G1,G2,j,2)
		benchmark_IsIn(n,G1,G2,j,2)
		benchmark_Pair(n,G1,G2,j,2)

	}
	fmt.Println("y2")
	fmt.Println(y2)

	f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
	f.SetCellValue("Sheet1","A1",nil)
	f.SetCellValue("Sheet1","B1","time Mexp G1 length n")
	f.SetCellValue("Sheet1","C1","time Mexp G2 length n")
	f.SetCellValue("Sheet1","D1","time conversion")
	f.SetCellValue("Sheet1","E1","time IsInSubGroup Affine length n")
	f.SetCellValue("Sheet1","F1","time IsInSubGroup Jac length n")
	f.SetCellValue("Sheet1","G1","time pairing length n")
	f.SetCellValue("Sheet1","H1","time IsInSubGroup Jac length 1")
	f.SetCellValue("Sheet1","I1","time pairing length 1")

    var i int
    
	for  i = 0; i < 7; i++ {
		
	f.SetCellValue("Sheet1","A"+strconv.Itoa(i+2),x1[i])
	f.SetCellValue("Sheet1","B"+strconv.Itoa(i+2),y1[i][0].Seconds())
	f.SetCellValue("Sheet1","C"+strconv.Itoa(i+2),y1[i][1].Seconds())
	f.SetCellValue("Sheet1","D"+strconv.Itoa(i+2),y1[i][2].Seconds())
	f.SetCellValue("Sheet1","E"+strconv.Itoa(i+2),y1[i][3].Seconds())
	f.SetCellValue("Sheet1","F"+strconv.Itoa(i+2),y1[i][4].Seconds())
	f.SetCellValue("Sheet1","G"+strconv.Itoa(i+2),y1[i][5].Seconds())
	f.SetCellValue("Sheet1","H"+strconv.Itoa(i+2),y1[i][6].Seconds())
	f.SetCellValue("Sheet1","I"+strconv.Itoa(i+2),y1[i][7].Seconds())
		
	 }





	 //-*---------------------------------------------------------

     f.NewSheet("Sheet2")


	 f.SetCellValue("Sheet2","A1",nil)
	f.SetCellValue("Sheet2","B1","time Mexp G1 length n")
	f.SetCellValue("Sheet2","C1","time Mexp G2 length n")
	f.SetCellValue("Sheet2","D1","time conversion")
	f.SetCellValue("Sheet2","E1","time IsInSubGroup Affine length n")
	f.SetCellValue("Sheet2","F1","time IsInSubGroup Jac length n")
	f.SetCellValue("Sheet2","G1","time pairing length n")
	f.SetCellValue("Sheet2","H1","time IsInSubGroup Jac length 1")
	f.SetCellValue("Sheet2","I1","time pairing length 1")

    
	for  i = 0; i < 6; i++ {
		
	f.SetCellValue("Sheet2","A"+strconv.Itoa(i+2),x2[i])
	f.SetCellValue("Sheet2","B"+strconv.Itoa(i+2),y2[i][0].Seconds())
	f.SetCellValue("Sheet2","C"+strconv.Itoa(i+2),y2[i][1].Seconds())
	f.SetCellValue("Sheet2","D"+strconv.Itoa(i+2),y2[i][2].Seconds())
	f.SetCellValue("Sheet2","E"+strconv.Itoa(i+2),y2[i][3].Seconds())
	f.SetCellValue("Sheet2","F"+strconv.Itoa(i+2),y2[i][4].Seconds())
	f.SetCellValue("Sheet2","G"+strconv.Itoa(i+2),y2[i][5].Seconds())
	f.SetCellValue("Sheet2","H"+strconv.Itoa(i+2),y2[i][6].Seconds())
	f.SetCellValue("Sheet2","I"+strconv.Itoa(i+2),y2[i][7].Seconds())
		
	 }
	 //f.SetCellValue("Sheet1","H"+strconv.Itoa(1),"time Mexp G1 length n")

	//var t=[]string {"rggr","fefe"}
	//f.SetCellValue(nil,"time Mexp G1 length n","time Mexp G2 length n","time conversion","time IsInSubGroup Affine length n","time IsInSubGroup Jac length n","time pairing length n","time IsInSubGroup Jac length 1","time pairing length 1")
	

	//f.SetCellValue(nil,"time Mexp G1 length n","time Mexp G2 length n","time conversion","time IsInSubGroup Affine length n","time IsInSubGroup Jac length n","time pairing length n","time IsInSubGroup Jac length 1","time pairing length 1")
	//f.SetCellValue(nil,"time Mexp G1 length n","time Mexp G2 length n","time conversion","time IsInSubGroup Affine length n","time IsInSubGroup Jac length n","time pairing length n","time IsInSubGroup Jac length 1","time pairing length 1")

	if err := f.SaveAs("basic_benchmarks_bls12-381.xlsx"); err != nil {
        fmt.Println(err)
    }
	//for j, n := range x2 {
	//	fmt.Println("n=",n)
	//benchmark_ME(n,G1,G2,j,2)
	//benchmark_IsIn(n,G1,G2,j,2)
	//benchmark_Pair(n,G1,G2,j,2)

//	}
	//fmt.Println("y2")
//	fmt.Println(y2)

//	t:=64
//	l:=7
//	n:=2;
//	fmt.Println("n=",n)
//	_,_,G1,G2:=curve.Generators()

//	benchmark_ME(n,G1,G2)
//	benchmark_Pair(n,G1,G2)
//	benchmark_ME(n,G1,G2)
//	benchmark_Pair(n,G1,G2)
//	benchmark_IsIn(100,G1,G2)
//	benchmark_eddsaverif(n)
//	benchmark_hash(n,G1)
//	benchmark_mult_Fq(n)
//	benchmark_IsOnCurve(n,G1)
//	sizes(1,G1,G2)
}


//Creates random vector with N entries in fr
func create_vectorFr(N int)[]fr.Element {
	c:=make([]fr.Element, N)
	for i:=range(c){
		(&c[i]).SetRandom()
	}
	return c		
}

//Creates random vector with N entries in big.Int
func create_vectorBig(N int)[]big.Int {
	c:=make([]big.Int, N)
	for i:=range(c){
		randInt, _ := crand.Int(crand.Reader, fr.Modulus())
		(&c[i]).Set(randInt)
	}
	return c		
}

//Creates vector with N entries in G1Affine by multiplying generator G1 of G1Affine with random coefficients
func create_vectorG1(N int, G1 curve.G1Affine)[]curve.G1Affine {
	coeffs:=create_vectorFr(N) //vector with N entries in Fr
	return curve.BatchScalarMultiplicationG1(&G1,coeffs) //returns (coeffs[i].G) for i=1...2nl
}

//Creates vector with N entries in G2Affine by multiplying generator G2 of G2Affine with random coefficients
func create_vectorG2(N int, G2 curve.G2Affine)[]curve.G2Affine {
	coeffs:=create_vectorFr(N) //vector with N entries in Fr
	return curve.BatchScalarMultiplicationG2(&G2,coeffs) //returns (coeffs[i].G) for i=1...2nl
}


//time for a MultiExp of size n
func benchmark_ME(n int, G1 curve.G1Affine, G2 curve.G2Affine, j int,x int){
	a:=create_vectorFr(n)
	var sum1 time.Duration
	var sum2 time.Duration	
	for i := 0; i < 100; i++ {
		start:=time.Now()
		A1:=create_vectorG1(n,G1)
//		fmt.Println("point vector created in", time.Since(start))
//		A2:=create_vectorG2(n,G2)
		c:=ecc.MultiExpConfig{}
		var X1 curve.G1Affine
//		var X2 curve.G2Affine
		start=time.Now()
		(&X1).MultiExp(A1,a,c)
		sum1+= time.Since(start)
//		start=time.Now()
//		(&X2).MultiExp(A2,a,c)
		sum2 += time.Since(start)
	}
	if(x==1){
		y1[j][0]=sum1/100
		y1[j][1]=sum2/100
	}else{
		y2[j][0]=sum1/100
		y2[j][1]=sum2/100

	}


//	fmt.Println("time Mexp G1 length n", sum1/100)
//	fmt.Println("time Mexp G2 length n", sum2/100)
}


// Time needed to check if n points lie on the curve
func benchmark_IsIn(n int, G1 curve.G1Affine, G2 curve.G2Affine, j int,x int){
	A:=create_vectorG1(n,G1)
	B:=make([]curve.G1Jac,n)
	st:=time.Now()
	for i:=range(B){
		(&B[i]).FromAffine(&A[i])
	}
	fmt.Println("time conversion", time.Since(st))
	if(x==1){
		y1[j][2]=time.Since(st)
	}else{
		y2[j][2]=time.Since(st)
	}
	st=time.Now()
	for i:=range(A){
		(&A[i]).IsInSubGroup()
	}
	if(x==1){
		y1[j][3]=time.Since(st)
	}else{
		y2[j][3]=time.Since(st)
	}	
//	fmt.Println("time IsInSubGroup Affine length n", time.Since(st))
	st=time.Now()
	for i:=range(B){
		(&B[i]).IsInSubGroup()
	}
	if(x==1){
		y1[j][4]=time.Since(st)
	}else{
		y2[j][4]=time.Since(st)
	}	
//	fmt.Println("time IsInSubGroup Jac length n", time.Since(st))
	st=time.Now()
	(&B[0]).IsInSubGroup()
	if(x==1){
		y1[j][6]=time.Since(st)
	}else{
		y2[j][6]=time.Since(st)

	}	
//	fmt.Println("time IsInSubGroup Jac length 1", time.Since(st))

}

// Time needed to compute n pairings
func benchmark_Pair(n int, G1 curve.G1Affine, G2 curve.G2Affine, j int,x int){
	var sum1, sum2 time.Duration
	for i := 0; i < 100; i++ {
	A:=create_vectorG1(n,G1)
	B:=create_vectorG2(n,G2)
	st:=time.Now()
	curve.Pair(A,B)
	sum1 += time.Since(st)

	st=time.Now()
	curve.Pair([]curve.G1Affine{A[0]},[]curve.G2Affine{B[0]})
	sum2 += time.Since(st)
	}
	if(x==1){
		y1[j][5]=sum1/100
//		fmt.Println("time pairing length n", sum1/100)
		y1[j][7]=sum2/100
//		fmt.Println("time pairing length 1", sum2/100)
	}else{
		y2[j][5]=sum1/100
//		fmt.Println("time pairing length n", sum1/100)
		y2[j][7]=sum2/100
//		fmt.Println("time pairing length 1", sum2/100)

	}	
}


// Sizes of points
func sizes(n int, G1 curve.G1Affine, G2 curve.G2Affine){
	A:=create_vectorG1(n,G1)
	B:=create_vectorG2(n,G2)
	var AJ curve.G1Jac
	var BJ curve.G2Jac
	(&AJ).FromAffine(&A[0])
	(&BJ).FromAffine(&B[0])
	//fmt.Println("theoretical size of G1", unsafe.Sizeof(A[0]))
//	fmt.Println("theoretical size of G2", unsafe.Sizeof(B[0]))
	//fmt.Println("real size of G1", size.Of(A[0]))
	//fmt.Println("real size of  G2", size.Of(B[0]))
	//fmt.Println("real size of G1 en Jac", size.Of(AJ))
	//fmt.Println("real size of G2 en Jac", size.Of(BJ))
}




// Computes eddsa signature of msg
func eddsasig(msg []byte)(eddsa.PublicKey, []byte){
	hFunc:=sha256.New()
	var signatureKey *eddsa.PrivateKey
	var signaturePublicKey eddsa.PublicKey
	signatureKey, _ = eddsa.GenerateKey(crand.Reader)
	signaturePublicKey = signatureKey.PublicKey
	signature, _ := signatureKey.Sign(msg, hFunc)
	return signaturePublicKey, signature
}


// Checks eddsa signature sig of msg with public key pk
func eddsaverif(msg []byte, pk eddsa.PublicKey, sig []byte)time.Duration{
	hFunc:=sha256.New()
	start := time.Now()
	_, _ = pk.Verify(sig, msg, hFunc)
	return time.Since(start)
}


// Checks n eddsa signatures
func benchmark_eddsaverif(n int){
	m, _ := crand.Int(crand.Reader, fr.Modulus())
	var bigm big.Int
	(&bigm).Set(m)
	msg:=(&bigm).Bytes()
	var total time.Duration
	for i:=0; i<n; i++{
	pk,sig:=eddsasig(msg)	
	t:=eddsaverif(msg, pk, sig)
	total+=t
	}
//	fmt.Println("time for n edDSA checks", total)
}



//measures time of n hashes of two points and a message in Fr
func benchmark_hash(n int, G1 curve.G1Affine) (bool, error){
	var total time.Duration
	hFunc:=sha256.New()
	m, _ := crand.Int(crand.Reader, fr.Modulus())
	var bigm big.Int
	(&bigm).Set(m)
	msg:=(&bigm).Bytes()
	for i:=0;i<n;i++{
		V:=create_vectorG1(2,G1);
		R:=V[0]
		A:=V[1]
		start:=time.Now()
		hFunc.Reset()
		sigRX := R.X.Bytes()
		sigRY := R.Y.Bytes()
		sigAX := A.X.Bytes()
		sigAY := A.Y.Bytes()
		toWrite := [][]byte{sigRX[:], sigRY[:], sigAX[:], sigAY[:], msg}
		for _, bytes := range toWrite {
			if _, err := hFunc.Write(bytes); err != nil {
				return false, err
			}
		}
		var hramInt big.Int
		hramBin := hFunc.Sum(nil)
		hramInt.SetBytes(hramBin)
		total += time.Since(start)
	}
//	fmt.Println("time to hash n*(2 points + message)", total)
	return true, nil
}

//measures time of checking if 2*n points are on curve
func benchmark_IsOnCurve(n int, G1 curve.G1Affine){
	V:=create_vectorG1(2*n,G1)
	start:=time.Now()
	for i:=0; i<2*n; i++{
		V[i].IsOnCurve()
	}
	fmt.Println("time to check if 2*n points are on curve", time.Since(start))
}



func benchmark_mult_Fq (n int){
	A:=make([]fr.Element, 3*n)
	B:=make([]fr.Element, 3*n)
	var total time.Duration
	var a,b fr.Element
	for j:=0;j<200; j++{
		for i:=0; i<3*n; i++{
			(&a).SetRandom()
			(&b).SetRandom()
			A[i]=a
			B[i]=b
		}

		start:=time.Now()
		for i:=0; i<3*n; i++{
			(&a).Mul(&A[i],&B[i])
		}
		total+= time.Since(start)
	}
//	fmt.Println("time for 3*n multiplications in Fr", total/200)
}


