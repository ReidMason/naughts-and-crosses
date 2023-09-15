import { type NextPage } from "next";
import Head from "next/head";
import { Registration } from "~/components/registration/Registration";

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Naughts and crosses</title>
      </Head>
      <main className="bg-slate-200 min-h-screen">
        <div className="flex md:items-center justify-center w-full h-screen">
          <Registration />
        </div>
      </main>
    </>
  );
};

export default Home;
