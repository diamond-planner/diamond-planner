import type { AuthProviderInfo, RecordService } from "pocketbase"

export async function providerLogin(
  provider: AuthProviderInfo,
  authCollection: RecordService,
  signup_key = ""
) {
  return await authCollection.authWithOAuth2({
    provider: provider.name,
    createData: {
      signup_key: signup_key,
    },
    query: {
      expand: "club",
    },
  })
}
