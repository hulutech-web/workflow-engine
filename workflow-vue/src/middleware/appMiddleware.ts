export default async () => {
  await Promise.all([useConfigStore().getConfig()])
}
