<script lang="ts">
  import {X} from "lucide-svelte";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";

  let {
    name = "",
    value = [] as string[],
    placeholder = "",
    validate = (args: { inputValue: string; value: string[] }) => true,
    onValueChange = (e: { value: string[] }) => {
    },
  } = $props();

  let inputValue = $state("");

  function addTag(raw: string) {
    const newTag = raw.trim();

    if (!newTag) {
      return;
    }

    if (!validate({inputValue: newTag, value})) {
      toastController.trigger({
        message: "You have entered a value that is not allowed here.",
        background: "preset-filled-error-500",
      });
      return;
    }

    if (value.includes(newTag)) {
      return;
    } // prevent duplicates

    const newValue = [...value, newTag];
    value = newValue;
    onValueChange({value: newValue});
    inputValue = "";
  }

  function removeTag(tag: string) {
    const newValue = value.filter((t) => t !== tag);
    value = newValue;
    onValueChange({value: newValue});
  }

  function handleKeyDown(e: KeyboardEvent) {
    if (e.key === "Enter" || e.key === ",") {
      e.preventDefault();
      addTag(inputValue);
    } else if (e.key === "Backspace" && !inputValue && value.length > 0) {
      removeTag(value[value.length - 1]);
    }
  }
</script>

<div class="tags-input input">
  <input
          name={name}
          oninput={(e) => (inputValue = (e.currentTarget.value))}
          onkeydown={handleKeyDown}
          placeholder={placeholder}
          type="text"
          value={inputValue}
  />

  <div class="tags-container">
    {#each value as tag (tag)}
    <span class="tag chip preset-filled-primary-500">
      {tag}
      <button
              type="button"
              class="remove-btn"
              aria-label="Remove tag"
              onclick={() => removeTag(tag)}
      >
        <X size={14}/>
      </button>
    </span>
    {/each}
  </div>
</div>

<style>
    .tags-input {
        border-radius: var(--radius-container);
    }

    .tags-container {
        display: flex;
        flex-wrap: wrap;
        gap: calc(var(--spacing) * 2);
        margin-block: 0.5em;
    }

    .tag {
        display: flex;
        align-items: center;
    }

    .remove-btn {
        background: none;
        border: none;
        cursor: pointer;
        line-height: 1;
    }

    input {
        flex: 1;
        border: none;
        outline: none;
        width: 100%;
        background-color: transparent;
        padding-inline: 0;
    }
</style>
