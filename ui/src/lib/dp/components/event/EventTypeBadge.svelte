<script lang="ts">
interface props {
  type: string
}

const { type }: props = $props()

function getDisplayString(): string {
  switch (type) {
    case "game":
      return "Game"
    case "practice":
      return "Practice"
    default:
      return "Event"
  }
}
</script>

<span
        class="badge"
        class:preset-filled-primary-500={type === "game"}
        class:preset-filled-secondary-500={type === "practice"}
        class:preset-filled-tertiary-500={type === "misc"}
>
{getDisplayString()}
</span>
