import { UsernameInput } from "../usernameInput/UsernameInput"

export const Registration = () => {
  return (
    <div className="flex flex-col justify-center md:grid md:grid-cols-2 p-8 md:p-24 items-center md:gap-16 md:rounded-lg max-w-7xl backdrop-blur-3xl overflow-hidden bg-black/20">
      {/* Remove this at some point       */}
      <div className="absolute -z-10 blur-2xl w-full h-full"
        style={{ backgroundImage: "url(https://cdn1.dotesports.com/wp-content/uploads/2023/02/02001438/Gawr-Gura-Return-to-Streaming.png)" }}
      ></div>

      <div className="flex flex-col gap-4">
        <h2 className="font-medium text-4xl">Register to play</h2>
        <p>Enter a username then create a game or join a friend</p>
      </div>

      <div className="flex-col flex gap-2 mt-10">
        <UsernameInput />
        <small className="ml-4 text-indigo-900"><a href="/">Transfer account</a></small>
      </div>
    </div>
  )
}
