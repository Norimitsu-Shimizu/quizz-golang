import type { AppProps } from 'next/app'
import React, { useState, createContext, useEffect, useCallback } from "react";
import { useRouter } from "next/navigation";
import { User } from "@/interfaces/index";
// import { getPublicConfig } from "@/utils/config";

export const AuthContext = createContext<any>("");


export default function App({ Component, pageProps }: AppProps) {
  // const config = getPublicConfig();
  const [currentUser, setCurrentUser] = useState<User>({
    mail_address: "",
  });
  const [isSignedIn, setIsSignedIn] = useState<boolean>();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [dataFetch, setDataFetch] = useState<boolean>(false);
  const [reload, setReload] = useState<boolean>(false);


  const router = useRouter();

  return (
    <AuthContext.Provider
    value={{
      currentUser,
      setCurrentUser,
      isSignedIn,
      setIsSignedIn,
      isLoading,
      setIsLoading,
      setDataFetch,
      dataFetch,
      reload,
      setReload,
    }}
  >
   <Component {...pageProps} />
  </AuthContext.Provider>
  )
}
