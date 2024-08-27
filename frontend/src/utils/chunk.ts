export function chunk<T>(arr: T[] = [], size: number): T[][] {
  const chunked: T[][] = []
  for (let i = 0; arr.length > i + size; i += size) {
    chunked.push(arr.slice(i, i + size))
  }
  return chunked
}
