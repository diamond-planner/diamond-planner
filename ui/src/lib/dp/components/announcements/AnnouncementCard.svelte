<script lang="ts">
  import {MessageCircle} from "lucide-svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.ts";
  import type {CustomAuthModel, ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import AnnouncementCoreContent from "./AnnouncementCoreContent.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  const {announcement}: Props = $props();
  const commentCount = announcement?.expand?.comments_via_announcement?.length ?? 0;

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
      announcement.expand?.club?.admins.includes(authRecord.id) ||
      announcement.expand?.team?.admins.includes(authRecord.id)
  );
</script>

<article class="card p-4 preset-tonal-surface shadow-xl h-full">
  <div class="flex flex-col justify-between h-full">
    <div>
      <AnnouncementCoreContent {announcement} textClasses="line-clamp-5"/>
    </div>

    <div>
      <hr class="mt-6 mb-2"/>

      <footer class="mt-3 flex gap-2 justify-between">

        {#if canEdit}
          <div class="flex gap-2">
            <AnnouncementForm showLabel={false} announcement={announcement} club={announcement.expand?.club ?? null}
                              team={announcement.expand?.team ?? null}
                              buttonClasses="btn btn-icon btn-sm preset-tonal-tertiary border border-tertiary-500"
            />

            <DeleteButton
                    id={announcement.id}
                    modelName="Announcement"
                    action={deleteAction}
                    classes="btn btn-icon btn-sm preset-tonal-error border border-error-500"
                    data-testid="delete-announcement-button"
            />
          </div>
        {:else }
          <div aria-hidden="true"></div>
        {/if}

        <div class="flex gap-4">
          {#if commentCount > 0}
            <div class="relative inline-block">
              <span class="sr-only">Number of Comments:</span>
              <span class="badge-icon preset-filled-primary-500 absolute z-10">
                {commentCount}
              </span>
              <MessageCircle aria-hidden="true"/>
            </div>
          {/if}

          <a
                  class="btn btn-sm preset-filled-primary-500"
                  href="/account/announcements/{announcement.id}">Read more</a
          >
        </div>
      </footer>
    </div>
  </div>
</article>

<style>
    .badge-icon {
        top: -9px;
        right: -6px;
        padding: 0.2rem;
    }
</style>