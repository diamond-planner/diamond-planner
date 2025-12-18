<script lang="ts">
// formats a date in PocketBase string format to local date for display

interface Props {
  date: string
}

let { date }: Props = $props()

const formattedDate = $derived(new Date(date))
</script>

{formattedDate.toLocaleString()}
