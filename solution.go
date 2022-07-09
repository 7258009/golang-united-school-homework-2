package square
import(
"math"
)

type YourTypeNameHere int
	
const(
	SidesCircle YourTypeNameHere = 0
	SidesTriangle YourTypeNameHere = 3
	SidesSquare YourTypeNameHere = 4
)

func CalcSquare(sideLen float64, sidesNum YourTypeNameHere) float64 {

	if (sidesNum == SidesCircle){
		return ( math.Pi * sideLen* sideLen)
			//(sideLen/(2*math.Pi)/2) )
		//(pi*r2)/2 
		//( math.Pi * (sideLen/(2*math.Pi)) * (sideLen/(2*math.Pi)) )
	}

	if (sidesNum == SidesSquare){
		return ( (math.Sqrt(3) * sideLen * sideLen) / 16 )
		//( (sideLen/4) * (sideLen/4) )
	}

	if (sidesNum == SidesTriangle){
		return (sideLen*math.Sqrt(sideLen*sideLen-sideLen*sideLen/4)/2)
	}
	
	return 0;
}

