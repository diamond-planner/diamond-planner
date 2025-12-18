<script lang="ts">
interface Props {
  background?: string
  classes?: string
  src?: string | undefined
  fallback?: string
}

let { background = "", classes = "", src, fallback = "" }: Props = $props()
</script>

{#if src}
  <figure
          class="avatar bg-surface-400-600 {classes} {background}"
          data-part="root"
          data-scope="avatar"
          data-testid="avatar"
  >
    <img
            alt="profile "
            class="avatar-image"
            data-part="image"
            data-scope="avatar"
            data-testid="avatar-image"
            {src}
    />
  </figure>
{:else}
  <span
          class="avatar-fallback {classes} {background}"
          data-part="fallback"
          data-scope="avatar"
          data-testid="avatar-fallback"
  >
    {fallback}</span>
{/if}
<style>
    .avatar {
        overflow: hidden;
        isolation: isolate;
        width: var(--size);
        border-radius: 100%;

        .avatar-image {
            width: 100%;
            object-fit: cover;
        }
    }

    .avatar-fallback {
        width: var(--size);
        height: var(--size);
        line-height: var(--size);
        text-align: center;
        display: inline-block;
        background-color: var(--color-primary-500);
        color: white;
        border-radius: 100%;
        font-weight: var(--font-weight-semibold);
    }
</style>
