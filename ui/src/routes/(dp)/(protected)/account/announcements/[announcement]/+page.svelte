<script lang="ts">
  import {Calendar, Clock, User} from "lucide-svelte";
  import PriorityBadge from "$lib/dp/components/announcements/PriorityBadge.svelte";
  import CommentSection from "$lib/dp/components/comments/CommentSection.svelte";
  import AnnouncementForm from "$lib/dp/components/forms/AnnouncementForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import type {PageProps} from "../../../../../../../.svelte-kit/types/src/routes";

  const {data}: PageProps = $props();
  const announcement = $derived(data.announcement);

  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const canEdit = $derived(
      $announcement.expand?.club?.admins.includes(authRecord.id) ||
      $announcement.expand?.team?.admins.includes(authRecord.id)
  );

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
    history.back();
  }

  const updated = $derived($announcement.updated);
  //@ts-expect-error - the multi-level expanding trips the typedef up
  const club = $derived($announcement?.expand?.club ?? $announcement?.expand?.team?.expand?.club);
</script>

<div class="flex justify-between items-center gap-4 lg:gap-6">
  <h1 class="h1 my-3!">{$announcement.title}</h1>
  <PriorityBadge priority={$announcement.priority}/>
</div>

<div class="flex justify-center">
  <div>
    <section class="my-6 space-y-3">
      <dl class="flex gap-2 items-center">
        {#if $announcement.expand?.team}
          <dt>Team:</dt>
          <dd class="badge preset-tonal-primary border border-primary-500">
            {$announcement.expand?.team?.name}
          </dd>
        {/if}

        {#if $announcement.expand?.club}
          <dt>Club:</dt>
          <dd class="badge preset-tonal-primary border border-primary-500">
            {$announcement.expand?.club?.name}
          </dd>
        {/if}
      </dl>

      <p class="flex gap-3 items-center">
        <User size="28"/>
        <span class="h6">{$announcement.expand?.author.first_name} {$announcement.expand?.author.last_name}</span>
      </p>

      <p class="flex gap-3 items-center">
        <Calendar size="28"/>
        <time class="h6" datetime="{updated}">{new Date(updated).toLocaleDateString()}</time>
      </p>

      <p class="flex gap-3 items-center">
        <Clock size="28"/>
        <time class="h6" datetime="{updated}">{new Date(updated).toLocaleTimeString()}</time>
      </p>
    </section>

    <section class="prose text-justify">
      {@html $announcement.bodytext}
    </section>

    {#if $announcement.link}
      <section class="my-6">
        <header>
          <h2 class="h4">Link</h2>
        </header>

        <a class="anchor block mt-2" href={$announcement.link} target="_blank"
           title="Link associated with this announcement: {$announcement.link}">
          {#if $announcement.link_text}
            {$announcement.link_text}
          {:else }
            {$announcement.link}
          {/if}
        </a>
      </section>
    {/if}

    <section class="my-6">
      <header>
        <h2 class="h4">Comments</h2>
      </header>
      <div class="mt-4 p-3 md:p-4 border border-surface-900-100 rounded-base">
        <CommentSection
                club={club}
                comments={$announcement?.expand?.comments_via_announcement ?? []}
                targetID={$announcement.id}
                targetType="announcement"
        />
      </div>
    </section>

    {#if canEdit}
      <hr class="my-6"/>

      <section class="my-8 card p-4 border border-surface-900-100">

        <header>
          <h2 class="h4">Admin Options</h2>
        </header>
        <div class="flex gap-3 mt-3 items-center flex-wrap">
          <AnnouncementForm showLabel={true} announcement={$announcement} club={$announcement.expand?.club ?? null}
                            team={$announcement.expand?.team ?? null}
                            buttonClasses="btn preset-tonal-tertiary border border-tertiary-500"/>

          <DeleteButton
                  id={$announcement.id}
                  modelName="Announcement"
                  action={deleteAction}
                  buttonText="Delete Announcement"
                  classes="btn preset-tonal-error border border-error-500"
                  data-testid="delete-announcement-button"
          />
        </div>
      </section>
    {/if}
  </div>
</div>
