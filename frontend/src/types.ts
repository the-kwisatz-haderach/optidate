export type CalendarEntry = {
  isSqueeze: boolean
  isHoliday: boolean
  isWeekend: boolean
  holidayName: string
  timestamp: string
  holidayTypes: string[] | null
  gapSize: number
  week: number
}

export type Country = {
  countryCode: string
  name: string
}
