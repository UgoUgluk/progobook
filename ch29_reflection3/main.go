package main

import (
	"reflect"
)

func inspectChannel(channel interface{}) {
	channelType := reflect.TypeOf(channel)
	if channelType.Kind() == reflect.Chan {
		Printfln("Type %v, Direction: %v",
			channelType.Elem(), channelType.ChanDir())
	}
}
func sendOverChannel(channel interface{}, data interface{}) {
	channelVal := reflect.ValueOf(channel)
	dataVal := reflect.ValueOf(data)
	if channelVal.Kind() == reflect.Chan &&
		dataVal.Kind() == reflect.Slice &&
		channelVal.Type().Elem() ==
			dataVal.Type().Elem() {
		for i := 0; i < dataVal.Len(); i++ {
			val := dataVal.Index(i)
			channelVal.Send(val)
		}
		channelVal.Close()
	} else {
		Printfln("Unexpected types: %v, %v",
			channelVal.Type(), dataVal.Type())
	}
}

func createChannelAndSend(data interface{}) interface{} {
	dataVal := reflect.ValueOf(data)
	channelType := reflect.ChanOf(reflect.BothDir,
		dataVal.Type().Elem())
	channel := reflect.MakeChan(channelType, 1)
	go func() {
		for i := 0; i < dataVal.Len(); i++ {
			channel.Send(dataVal.Index(i))
		}
		channel.Close()
	}()
	return channel.Interface()
}
func readChannels(channels ...interface{}) {
	channelsVal := reflect.ValueOf(channels)
	cases := []reflect.SelectCase{}
	for i := 0; i < channelsVal.Len(); i++ {
		cases = append(cases, reflect.SelectCase{
			Chan: channelsVal.Index(i).Elem(),
			Dir:  reflect.SelectRecv,
		})
	}
	for {
		caseIndex, val, ok := reflect.Select(cases)
		if ok {
			Printfln("Value read: %v, Type: %v", val,
				val.Type())
		} else {
			if len(cases) == 1 {
				Printfln("All channels closed.")
				return
			}
			cases = append(cases[:caseIndex],
				cases[caseIndex+1:]...)
		}
	}
}

func main() {
	//inspectChannel
	var c chan<- string
	inspectChannel(c)

	//sendOverChannel
	values1 := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel1 := make(chan string)
	go sendOverChannel(channel1, values1)
	for {
		if val, open := <-channel1; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}

	//createChannelAndSend
	values := []string{"Alice", "Bob", "Charlie", "Dora"}
	channel := createChannelAndSend(values).(chan string)
	for {
		if val, open := <-channel; open {
			Printfln("Received value: %v", val)
		} else {
			break
		}
	}

	//readChannels
	cities := []string{"London", "Rome", "Paris"}
	cityChannel := createChannelAndSend(cities).(chan string)
	prices := []float64{279, 48.95, 19.50}
	priceChannel := createChannelAndSend(prices).(chan float64)
	readChannels(channel, cityChannel, priceChannel)

}
