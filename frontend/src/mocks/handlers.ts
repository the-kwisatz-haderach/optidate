import { http, HttpResponse, passthrough } from 'msw'
import { calendarStub } from './calendarStub'
import { countriesStub } from './countriesStub'

export const handlers = [
  http.get('/calendar/*', () => HttpResponse.json(calendarStub)),
  http.get('/countries', () => HttpResponse.json(countriesStub)),
  http.get('*', () => {
    return passthrough()
  }),
]
