package library

const secondsInMinute = 60
const minutesInHour = 60
const HourInADay = 24

func daysToHour(day int) int {
	return day * HourInADay
}

func DaysToMinutes(day int) int {
	return day * HourInADay * minutesInHour
}
