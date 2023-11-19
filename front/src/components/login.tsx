"use client";

import { useForm, SubmitHandler } from "react-hook-form";
import { SignInParams } from "@/interfaces";
import { signIn } from "@/api/auth";
import { AuthContext } from "@/store/authcontext";
import React, { useState, useContext } from "react";
import { useRouter } from "next/navigation";
import Button from "@/atoms/button";

const LoginFormSection = () => {
  const { currentUser, setCurrentUser, setIsSignedIn } = useContext(AuthContext);

  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    handleSubmit,
    setError,
    clearErrors,
    formState: { isValid, errors },
  } = useForm<{
    email: string;
    password: string;
    auth: boolean;
  }>({
    mode: "onBlur",
  });

  const router = useRouter();

  const onSubmit: SubmitHandler<SignInParams> = async (data: SignInParams) => {
    const loginData = {
      ...data,
    };

    setIsLoading(true);

    try {
      const res = await signIn(loginData);

      if (res.status === 200) {
        // localStorage.setItem("_access_token", res.headers["access-token"]);
        // localStorage.setItem("_client", res.headers["client"]);
        // localStorage.setItem("_uid", res.headers["uid"]);
        const token = res.data.user.token;
        // tokenをクライアント側で使えるように保存する
        localStorage.setItem('JWT_TOKEN', token);

        console.log(res.data.user)
        const user = res.data.user
        // setCurrentUser({
        //   username: user.username,
        //   email: user.email,
        // });
        setCurrentUser(user);
        console.log(currentUser)
        setIsSignedIn(true);
        router.push(`/`);
      } else {
        setError("auth", {
          type: "custom",
          message: "メールアドレスまたはパスワードが正しくありません。",
        });
        setIsLoading(false);
      }
    } catch (e) {
      setError("auth", {
        type: "custom",
        message: "メールアドレスまたはパスワードが正しくありません。",
      });
      setIsLoading(false);
    }
  };

  return (
    <>
      <div className="mx-auto max-w-[424px] rounded-20 py-6 flex flex-col items-center px-3 lg:h-full lg:rounded lg:min-h-screen pb-10 lg:pb-0">
        <div className="text-text-color text-title lg:text-xl font-bold mt-4 mb-6 sm:mb-8">
          {/* eslint-disable-next-line @next/next/no-img-element */}

          <h1 className="block w-max mx-auto text-title sm:text-lg">
            ログイン
          </h1>
        </div>

        <form className="text-center w-full" onSubmit={handleSubmit(onSubmit)}>
          <div className="flex flex-col mx-auto">
            <label className="text-text-color text-left mb-1.5 font-semibold text-base">
              メールアドレス
            </label>
            <input
              className="indent-3 border-[1px] border-divider-gray h-[45px] mb-2 rounded-10 placeholder:text-base placeholder:text-gray"
              id="Email"
              type="text"
              placeholder="test@example.com"
              {...register("email", {
                required: {
                  value: true,
                  message: "メールアドレスを入力して下さい。",
                },
                onChange: () => clearErrors("auth"),
              })}
            />
            <div className="text-error-color text-left text-base">
              {errors.email && errors.email.message}
            </div>
          </div>
          <div className="flex flex-col mx-auto mb-4">
            <label className="text-text-color text-left mt-3 mb-1.5 font-semibold text-base">
              パスワード
            </label>
            <input
              className="indent-3 border-[1px] border-divider-gray h-[45px] mb-1 rounded-10 placeholder:text-base"
              type="password"
              id="password"
              {...register("password", {
                required: {
                  value: true,
                  message: "パスワードを入力して下さい。",
                },
                onChange: () => clearErrors("auth"),
              })}
            />
            <div className="text-error-color text-left text-base">
              {errors.password && errors.password.message}
            </div>
          </div>

          {errors["auth"] && (
            <p className="text-base text-error-color mb-4 tracking-tighter text-left">
              メールアドレスまたはパスワード<br className="sm:hidden"></br>
              が正しくありません。
            </p>
          )}

          <Button
            color={"purple"}
            disabled={!isValid}
            type="submit"
            fullWidth={true}
          >
            {isLoading ? (
              <div className="w-max flex py-[8.5px] mx-auto">
                <div className="animate-ping h-1 w-1 bg-background rounded-full"></div>
                <div className="animate-ping h-1 w-1 bg-background rounded-full mx-[10px]"></div>
                <div className="animate-ping h-1 w-1 bg-background rounded-full"></div>
              </div>
            ) : (
              "次へ"
            )}
          </Button>

          <p className="text-dark-gray text-base mb-[32px] mt-[20px] font-light">
            パスワードを忘れた方は、
            <span
              onClick={() => router.push("/password-mail-input")}
              className="text-main-color cursor-pointer font-bold"
            >
              こちら
            </span>
          </p>
        </form>
      </div>
    </>
  );
}

export default LoginFormSection

