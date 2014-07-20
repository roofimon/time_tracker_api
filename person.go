package time_tracker

type Person struct {
    Name string
    Site string
    Checkin int64
    Checkout int64 "checkout"
}

