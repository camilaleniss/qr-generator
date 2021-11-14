package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBase64QRCode(t *testing.T) {
	c := require.New(t)

	encodedQR, err := GetBase64QRCode("Hola BB")
	c.Nil(err)

	expectedQR := "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABFElEQVR42uyYsXXDMAxE4aeCJUbQKBxNGU2jcASWLPx0eQBlWmZkuUxeeFfZ8q/wcOeDhKIo6q9LUbXGJDJtt/1rIdAD9eOyypwDthuOT4cCAoAK+KBkATKB90D0NSLwGditR+AdsHszIoe7XJh3dOAZ5lo+pP3QQJNNEttFT/jfgLZgShLudU6pXzkCtlGl+i0mRZ1kTJqFQAdIKNN2zChgnUHgDPBJRiQt/vtXZ04CjzB3R3qPmqwp2MNMoAeeaW89ygFzaRkNeD1pLaQWdCFG4PXmFfF6YNbT09t/dKBdKIrzZk6gfznwKJwzQOAS8En69gmBH0Dz5l4PrJkPCLQwP/z7n9/+YwMURVG/o+8AAAD//3GJQZDBLZ76AAAAAElFTkSuQmCC"

	c.Equal(expectedQR, encodedQR)
}
