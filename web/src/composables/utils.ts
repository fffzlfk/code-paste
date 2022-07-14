export function UTC2Local(date: string): string {
  const dateObj = new Date(date)
  const localDate = dateObj.toLocaleString([], { year: 'numeric', month: 'numeric', day: 'numeric', hour: '2-digit', minute: '2-digit' })
  return localDate
}
