import ky from "ky";

const BASE_URL = "http://localhost:8080/api/v2";

export const apiCustomToken = (token: string) =>
  ky.create({
    prefixUrl: BASE_URL,
    hooks: {
      beforeRequest: [
        async (request) => {
          request.headers.set("Authorization", `Bearer ${token}`);
        },
      ],
    },
  });
