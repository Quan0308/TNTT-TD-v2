import { signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "firebase-config";
import { apiCustomToken } from ".";

export const SignInService = async (email: string, password: string) => {
  try {
    const userCredential = await signInWithEmailAndPassword(
      auth,
      email,
      password,
    );
    const idToken = await userCredential.user.getIdToken();
    const accessToken = await apiCustomToken(idToken).post("auth/login");
    return accessToken;
  } catch (err) {
    console.error(err);
  }
};
