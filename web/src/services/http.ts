import ky from "ky"

export function fetchData(page: number, input: string): Promise<SearchResponse> {
  return ky.post("http://127.0.0.1:3000/search", {
    json: { page, input }
  }).json()
}