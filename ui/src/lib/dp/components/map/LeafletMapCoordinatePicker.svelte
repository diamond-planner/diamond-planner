<script lang="ts">
import { onMount } from "svelte"
import L, { type LeafletMouseEvent } from "leaflet"
import "leaflet/dist/leaflet.css"
import { env } from "$env/dynamic/public"

const DEFAULT_ZOOM = 13

interface Props {
  latitude: number
  longitude: number
}

let { latitude = $bindable(), longitude = $bindable() }: Props = $props()

let mapContainer: HTMLElement

onMount(() => {
  function onMapClick(e: LeafletMouseEvent) {
    markers.clearLayers()

    let marker = L.marker([e.latlng.lat, e.latlng.lng])
    markers.addLayer(marker)
    markers.addTo(map)

    latitude = e.latlng.lat
    longitude = e.latlng.lng
  }

  let markers = L.layerGroup()
  const map = L.map(mapContainer).setView([latitude, longitude], DEFAULT_ZOOM)

  L.tileLayer(`https://{s}.${env.PUBLIC_TILE_SERVER}/{z}/{x}/{y}.png`, {
    maxZoom: 19,
    detectRetina: true,
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map)

  map.on("click", onMapClick)

  $effect(() => {
    map.setView([latitude, longitude])
  })

  return () => {
    map.remove()
  }
})
</script>

<div bind:this={mapContainer} id="map" class=" shadow-xl rounded-base border-4 border-surface-500"></div>

<style>
    #map {
        height: 20rem;
        @media (min-width: 800px) {
            height: 30rem;
        }
    }
</style>

