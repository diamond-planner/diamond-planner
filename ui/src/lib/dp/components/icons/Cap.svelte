<script lang="ts">
interface Props {
  identifier: string // Skylarks Caps get assigned a unique identifier, for all others this is just a color
  classes: string
}

let { identifier, classes }: Props = $props()
</script>

{#if identifier === "B"}
  <img class="{classes}" src="/uniform/B_cap.svg" alt="Berlin Skylarks Home Cap">
{:else if identifier === "Bird"}
  <img class="{classes}" src="/uniform/Bird_cap.svg" alt="Berlin Skylarks Road Cap">
{:else}
  <div class="{classes}">
    <svg width="100%" height="100%" viewBox="0 0 800 600" xmlns="http://www.w3.org/2000/svg"
         xml:space="preserve"
         style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2;">
      <g id="path3056" transform="matrix(88.2013,0,0,88.2013,-242.384,4.88261)">
          <path
            d="M8.969,6.681C9.211,6.631 9.419,6.521 9.611,6.341C9.841,6.125 9.964,5.929 10.267,5.288C10.362,5.088 10.474,4.865 10.516,4.793L10.593,4.662L10.733,4.615C10.811,4.59 11.016,4.535 11.189,4.494C11.363,4.453 11.507,4.417 11.51,4.414C11.514,4.411 11.524,4.262 11.534,4.083C11.56,3.579 11.521,3.044 11.426,2.641C11.177,1.582 10.561,0.888 9.46,0.427C9.153,0.298 9.06,0.269 8.894,0.251C8.701,0.229 8.654,0.209 8.563,0.111C8.48,0.021 8.431,0 8.303,0C8.215,-0 8.059,0.025 8.04,0.042C8.033,0.048 8.023,0.09 8.018,0.135L8.009,0.216L7.831,0.226C7.013,0.271 6.408,0.426 5.966,0.701C5.532,0.972 5.273,1.302 5.05,1.867C4.978,2.049 4.949,2.137 4.752,2.765C4.701,2.926 4.624,3.146 4.581,3.255C4.469,3.535 4.442,3.644 4.442,3.804C4.441,3.993 4.417,4.037 4.18,4.268C4.076,4.369 3.818,4.641 3.607,4.871C3.397,5.101 3.126,5.384 3.007,5.499C2.808,5.692 2.787,5.717 2.765,5.79C2.743,5.862 2.742,5.955 2.765,5.975C2.769,5.979 2.804,5.971 2.843,5.957C2.895,5.939 2.954,5.894 3.066,5.789C3.198,5.665 3.234,5.64 3.329,5.602C3.621,5.488 4.152,5.428 4.582,5.463C5.3,5.52 5.895,5.704 7.049,6.226C7.618,6.484 7.889,6.585 8.208,6.657C8.41,6.702 8.806,6.715 8.969,6.681Z"
            style="fill:{identifier};fill-rule:nonzero;stroke:white;stroke-width:0.18px;"
          />
      </g>
    </svg>
  </div>
{/if}