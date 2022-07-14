package pkg

import "github.com/yu-hasebe/rock-paper-scissors/pb"

func EncodeHandShapes(n int32) pb.HandShapes {
	switch n {
	case 1:
		return pb.HandShapes_ROCK
	case 2:
		return pb.HandShapes_PAPER
	case 3:
		return pb.HandShapes_SCISSORS
	case 4:
		return pb.HandShapes_HAND_SHAPES_UNKNOWN
	}
}
