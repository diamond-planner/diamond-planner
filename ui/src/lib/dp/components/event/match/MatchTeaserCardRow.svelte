<script lang="ts">
interface Props {
  logo: string
  teamName: string
  score: number | undefined
  subduedScore: boolean
}

let { logo, teamName, score, subduedScore }: Props = $props()
</script>

<div class="flex justify-between items-center">
    <img src="{logo}" alt="team logo for {teamName}" width="50" loading="lazy"/>
    <span>{teamName}</span>
    <span class:text-surface-700-300={subduedScore} class="text-xl font-extrabold ">{score ?? " "}</span>
</div>