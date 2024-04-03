export default function SearchCard({ title, description, url, image }: {title: string, description: string, url: string, image: string}) {
  return (
    <div key={url} className="max-w-md md:max-w-2xl">
      <div className="flex items-center mb-4">
        <img className="h-10 w-10 rounded-full mr-3" src={image || "./src/assets/placeholder.png"} alt="logo" />
        <a href={url} className="block mt-1 text-lg leading-tight font-medium text-indigo-700 hover:underline line-clamp-1">{title}</a>
      </div>
      <p className="mt-2 text-xs text-gray-700 line-clamp-2">{description}</p>
    </div>
  )
}
