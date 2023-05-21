import { useForm, SubmitHandler } from "react-hook-form";
import { useRouter } from "next/router";
import { SignUpParams } from "../../interfaces";
import { signUp } from "../api/auth";
import { AuthContext } from "@/pages/_app";
import React, { useState, useContext } from "react";
import Button from "@/atoms/button";
import Link from "next/link";
import { getPublicConfig } from "@/utils/config";

const SignupFormSection = () => {
  const config = getPublicConfig();
  const { setCurrentUser, setMail } = useContext(AuthContext);
  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    setError,
    formState: { isValid, errors },
    handleSubmit,
  } = useForm<SignUpParams>({
    mode: "onBlur",
  });

  const router = useRouter();

  const onSubmit: SubmitHandler<SignUpParams> = async (data: SignUpParams) => {
    const postData = {
      ...data,
    };

    setIsLoading(true);

    try {
      const res = await signUp(postData);

      if (res.status === 200) {
        localStorage.setItem("_access_token", res.headers["access-token"]);
        localStorage.setItem("_client", res.headers["client"]);
        setCurrentUser(res.data.data);
        setMail(data.mail_address);
        router.push("/signup/mail-sent");
      } else {
        setError("mail_address", {
          type: "custom",
          message: "既に登録されているメールアドレスです。",
        });
        setIsLoading(false);
      }
    } catch (err) {
      setError("mail_address", {
        type: "custom",
        message: "既に登録されているメールアドレスです。",
      });
      setIsLoading(false);
    }
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <label>
            氏名
            <span>必須</span>
          </label>

          <input
            id="LastName"
            type="text"
            placeholder="test"
            {...register("last_name", {
              required: "名前を入力して下さい。",
            })}
          />
          <p>
            {errors.last_name && errors.last_name.message}
          </p>

          <input
            id="FirstName"
            type="text"
            placeholder="taro"
            {...register("first_name", {
              required: "名前を入力して下さい。",
            })}
          />
          <p>
            {errors.first_name && errors.first_name.message}
          </p>
        </div>

        <div>
          <label>
            メールアドレス
            <span>必須</span>
          </label>

          <input
            id="LastName"
            type="text"
            placeholder="test@example.com"
            {...register("mail_address", {
              required: "メールアドレスを入力して下さい。",
              pattern: {
                value: /^[\w\-._+]+@[\w\-._]+\.[A-Za-z]+/,
                message: "メールアドレスの形式が正しくありません。",
              },
            })}
          />
          <p>
            {errors.mail_address && errors.mail_address.message}
          </p>
        </div>
        <div className="flex flex-col mx-auto">
          <label>
            パスワード
            <span>必須</span>
          </label>
          <input
            type="password"
            id="password"
            {...register("password", {
              required: "パスワードを入力して下さい。",
              minLength: {
                value: 6,
                message: "6文字以上で入力してください。",
              },
              maxLength: {
                value: 128,
                message: "128文字以内で入力してください。",
              },
            })}
          />
          <p className="text-error-color text-left text-base tracking-tighter mt-2">
            {errors.password && errors.password.message}
          </p>
        </div>

        <Button
          color={"purple"}
          disabled={!isValid}
          type="submit"
          fullWidth={true}
        >
          送信
        </Button>
      </form>
    </>
  )
}

export default SignupFormSection
