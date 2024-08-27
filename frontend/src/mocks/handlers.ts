import { http, HttpResponse, passthrough } from 'msw'
import { calendarStubHR as calendarStub } from './calendarStub'
import { countriesStub } from './countriesStub'

export const handlers = [
  http.get('/calendar/*', () => HttpResponse.json(calendarStub)),
  http.get('/countries', () => HttpResponse.json(countriesStub)),
  http.get('https://flagsapi.com/*', async () => {
    const buffer = await fetch(`/flag.png`).then((response) =>
      response.arrayBuffer()
    )
    return HttpResponse.arrayBuffer(buffer, {
      headers: {
        'Content-Type': 'image/jpeg',
      },
    })
  }),
  http.get('*', () => {
    return passthrough()
  }),
]
