type Search = {
  page: number
  search: string
}

interface PageResult {
  Title: string
  Url: string
  Description: string
  Image: string
}

interface SearchResponse {
  pages: PageResult[]
  totalPages: number
}