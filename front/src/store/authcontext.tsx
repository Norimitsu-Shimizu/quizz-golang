"use client";

// import type { AppProps } from 'next/app'
import React, { useState, createContext, ReactNode } from "react";
// import { useRouter } from "next/navigation";
import { User } from "@/interfaces/index";
// import { getPublicConfig } from "@/utils/config";

export const AuthContext = createContext<any>("");

export const UserContext = ({ children }: { children: ReactNode }) => {
  // const config = getPublicConfig();
  const [currentUser, setCurrentUser] = useState<User>({
    username: "",
    email: "",
    // token: "",
  });
  const [isSignedIn, setIsSignedIn] = useState<boolean>();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [dataFetch, setDataFetch] = useState<boolean>(false);
  const [reload, setReload] = useState<boolean>(false);

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
      {children}
    </AuthContext.Provider>
  )
}
