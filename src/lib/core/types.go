package core

import "github.com/gizak/termui/v3/widgets"

type (
	////////////////////////////////////
	// Custom Datastructures

	WhlItem struct {
		Value string
		Next  *WhlItem
	}

	Wheel struct {
		First   *WhlItem
		Current *WhlItem
		Next    *Wheel
		Chars   *[]string
		Seeds   *Wheel
	}

	Combination struct {
		First Wheel
		Chars []string
		Seed  Wheel
	}

	////////////////////////////////////
	// Channels and similar

	// Signal type, for sending a signal to a channel
	Signal struct{}

	/* A simple signal with multible different Values
	   (?) 	In Most Usecases this is used in some sort
			of Control Channel (i.e. theres a code for stop, start, joining, joined etc.)

	   (?)	There will also be a Enum, where some Signal Codes are Predefined
	*/
	CodeSignal struct {
		Code int // The Signal Code
	}

	// For content rich Messages
	Message struct {
		Content any    // to attach values to Message
		Msg     string // The actual message
	}

	CH[T Signal | CodeSignal | Message] (chan T)

	////////////////////////////////////
	// Enum types

	HashMethod uint8
	SignalCode int

	///////////////////////////////////
	// Abstract Custom Data-Types (i.e. to group Data)

	ResultPair struct {
		PlainText string
		Hash      string
	}

	//////////////////////////////////
	// User Interface

	UiComponent interface {
		Render()
	}

	UiTxtComponent struct {
		Text  string
		Title string
		Ph    widgets.Paragraph
	}

	UiCrackerComponent struct {
		Height   int
		Seed     UiTxtComponent
		Hash     UiTxtComponent
		PlainTxt UiTxtComponent
	}

	UI struct {
		Components []*UiComponent
	}

	/////////////////////////////////
	// High Level Types (Functionality types)

	// The Type for Threads, wich crack hashes
	Cracker struct {
		FoundCH      chan ResultPair // Sends in this channel when the Plaintext has been found, Main-Channel Closes this channel, on wich all threads are listening, so they terminate if closed
		JoinCH       chan<- Signal   // Sends signal when Joined, this channel will be set to <thread-count>, means only when all threads send Joined, this channel holds value, is received by Main-Thread
		Combination  Combination
		SeedStr      string // The seeds whitch this Cracker uses
		Chars        string
		Ui_Component UiCrackerComponent
		HashMethod   HashMethod
	}
)

// HashMethod Enum-Values
const (
	SHA1   HashMethod = 0
	SHA256 HashMethod = 1
	SHA512 HashMethod = 2
	MD5    HashMethod = 3
)
