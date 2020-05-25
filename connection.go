package quickfix

import "io"

func writeLoop(connection io.Writer, messageOut chan []byte, log Log) {
	for {
		// The value of ok is true if the value received was delivered by a
		// successful send operation to the channel, or false if it is a zero
		// value generated because the channel is closed and empty.
		msg, ok := <-messageOut
		if !ok {
			log.OnEvent(logWithTrace("exiting connection writeloop"))
			return
		}

		if _, err := connection.Write(msg); err != nil {
			log.OnEvent(err.Error())
		}
	}
}

func readLoop(parser *parser, msgIn chan fixIn, log Log) {
	defer close(msgIn)

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.OnEvent(logWithTracef("read loop connection terminatied: %s", err.Error()))
			} else {
				log.OnEvent(logWithTracef("read loop error: %s", err.Error()))
			}
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
