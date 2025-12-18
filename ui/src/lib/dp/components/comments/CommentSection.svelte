<script lang="ts">
  import {Send} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import CommentRow from "$lib/dp/components/comments/CommentRow.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedComment} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, CommentsCreate, CommentsResponse} from "$lib/dp/types/pb-types.ts";

  interface Props {
    targetID: string;
    targetType: "announcement" | "event";
    comments: ExpandedComment[];
    club: ClubsResponse;
  }

  const {targetID, targetType, comments, club}: Props = $props();

  let commentText = $state("");

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  async function addComment(event: Event) {
    event.preventDefault();

    const formValues: CommentsCreate = {
      text: commentText,
      user: authRecord.id,
    };

    if (targetType === "event") {
      formValues.event = targetID;
    } else if (targetType === "announcement") {
      formValues.announcement = targetID;
    }

    try {
      await client.collection("comments").create<CommentsResponse>(formValues);
    } catch (error) {
      toastController.trigger({
        message: "Failed to add comment.",
        background: "preset-filled-error-500",
      });
    }

    commentText = "";
    await invalidate("comments:list");
  }
</script>

{#if comments.length > 0}
  <ul>
    {#each comments as comment (comment.id)}
      <li class={["my-3 md:my-4 flex gap-2", authRecord.id === comment?.expand?.user?.id && "flex-row-reverse"]}>
        <CommentRow {comment} {club}/>
      </li>
    {/each}
  </ul>
{:else}
  <p>No comments yet.</p>
{/if}

<form class="mt-6">
  <label class="label mb-2" for="new-comment-input">Comment</label>
  <div class="input-group grid-cols-[1fr_auto]">
    <input
            bind:value={commentText}
            class="ig-input rounded-s-base"
            id="new-comment-input"
            placeholder="Your comment..."
            type="text"
    />
    <button class="ig-btn preset-filled" disabled={!commentText} onclick={addComment} title="Add comment" type="submit">
      <Send size={18}/>
    </button>
  </div>
</form>
