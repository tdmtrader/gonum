package goblas

import (
	"testing"

	"github.com/gonum/blas"
	"github.com/gonum/blas/testblas"
)

func BenchmarkDgemmSmSmSm(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.SmallMat,
		testblas.SmallMat,
		testblas.SmallMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmMedMedMed(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.MediumMat,
		testblas.MediumMat,
		testblas.MediumMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmMedLgMed(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.MediumMat,
		testblas.LargeMat,
		testblas.MediumMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmLgLgLg(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.LargeMat,
		testblas.LargeMat,
		testblas.LargeMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmLgSmLg(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.LargeMat,
		testblas.SmallMat,
		testblas.LargeMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmLgLgSm(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.LargeMat,
		testblas.SmallMat,
		testblas.LargeMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmHgHgSm(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.HugeMat,
		testblas.SmallMat,
		testblas.HugeMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmMedMedMedTNT(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.MediumMat,
		testblas.MediumMat,
		testblas.MediumMat,
		blas.Trans,
		blas.NoTrans,
	)
}

func BenchmarkDgemmMedMedMedNTT(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.MediumMat,
		testblas.MediumMat,
		testblas.MediumMat,
		blas.NoTrans,
		blas.Trans,
	)
}

func BenchmarkDgemmMedMedMedNTNT(b *testing.B) {
	testblas.DgemmBenchmark(b,
		Blas{},
		testblas.MediumMat,
		testblas.MediumMat,
		testblas.MediumMat,
		blas.NoTrans,
		blas.NoTrans,
	)
}
