"use server";

export async function signInWithEmailAndPassword(data: {
  email: string;
  password: string;
}) {
  // const supabase = await createSupabaseServerClient();
  // const result = await supabase.auth.signInWithPassword({
  //   email: data.email,
  //   password: data.password,
  // });

  // return JSON.stringify(result);
  return JSON.stringify(data);
}
