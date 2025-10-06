# Page 4

## Text Content

```
<!-- Variants Table with Custom Checkboxes -->
<div class="bg-gray-200/70 mx-5

dark:bg-white/20 dark:text-stone-200 rounded-lg overfl

<div class="px-4 py-2 rounded-lg border-b">
<h5 class="text-lg font-light">Variants (<%= @product[:variant_list].size %>)</h5>
</div>
<% if @product[:variant_list].empty? %>
<div class="p-4 text-center rounded-lg text-gray-500">
<p>No variants found for this product.</p>
</div>
<% else %>
<div class="overflow-x-auto rounded-lg">
<table class="min-w-full shadow-sm rounded-lg dark:bg-white/20 dark:text-stone-20
<thead class="">
<tr class="text-center">
<!-- Header checkbox for selecting all variants -->
<th scope="col" class="dark:bg-white/20 dark:text-stone-200 relative px-4 sm:w
<div class="group relative grid size-4 grid-cols-1">
<input id="select-all-variants" type="checkbox" class="col-start-1 row-st
<svg class="pointer-events-none col-start-1 row-start-1 size-3.5 self-cen
<path class="opacity-0 group-[input:checked]:opacity-100" d="M3 8L6 11L
</svg>
</div>
</th>
<th class="dark:bg-white/10 dark:text-stone-200 px-4 py-2 text-sm font-light
<th class="px-4 py-2 text-sm font-light text-gray-500 dark:text-stone-200">SK
<th class="px-4 py-2 text-sm font-light text-gray-500 dark:text-stone-200">Inv
<th class="px-4 py-2 text-sm font-light text-gray-500 dark:text-stone-200">Op
<th class="px-4 py-2 text-sm font-light text-gray-500 dark:text-stone-200">Ac
</tr>
</thead>
<tbody class="bg-white text-center rounded-lg divide-y divide-gray-200">
<% @product[:variant_list].each do |variant| %>
<tr class="hover:bg-gray-50">
<!-- Checkbox for each variant -->
<td class="relative px-4 sm:w-12 sm:px-6">
<div class="group relative grid size-4 grid-cols-1">
<input type="checkbox" class="variant-checkbox col-start-1 row-start-1
<svg class="pointer-events-none col-start-1 row-start-1 size-3.5 self-c
<path class="opacity-0 group-[input:checked]:opacity-100" d="M3 8L6 1
</svg>
</div>
</td>
<td class="px-4 py-2 text-sm text-gray-700"><%= variant[:title] %></td>


```

