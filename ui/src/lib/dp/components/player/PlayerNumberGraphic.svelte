<script lang="ts">
interface Props {
  content: string
  classes?: string
}

let { content, classes = "preset-filled-primary-500" }: Props = $props()
</script>

<span class="text-4xl font-bold {classes}">
    {content}
</span>

<style lang="postcss">
    span {
        display: block;
        height: 4rem;
        width: 4rem;
        line-height: 4rem;
        border-radius: 50%;
        text-align: center;
    }
</style>