<script lang="ts">
  import LocalePicker from "$lib/dp/components/meta/LocalePicker.svelte";
  import LoginBadge from "$lib/dp/auth/LoginBadge.svelte";
</script>

<LocalePicker/>

<LoginBadge signupAllowed={true}/>