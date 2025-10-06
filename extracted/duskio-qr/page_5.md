# Page 5

## Text Content

```
<td class="px-4 py-2 text-sm text-gray-700">
<code class="bg-gray-100 px-1 rounded"><%= variant[:sku] %></code>
</td>
<td class="px-4 py-2 text-sm">
<% inventory = variant[:inventory_quantity].to_i %>
<div class="py-2 px-3 inline-block bg-white border border-gray-200 rounde
<div class="flex items-center gap-x-1.5">
<button type="button" class="size-6 inline-flex justify-center items<svg class="shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" w
<path d="M5 12h14"></path>
</svg>
</button>
<input class="p-1 w-12 bg-transparent border-0 text-gray-800 text-cen
<button type="button" class="size-6 inline-flex justify-center items<svg class="shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" w
<path d="M5 12h14"></path>
<path d="M12 5v14"></path>
</svg>
</button>
</div>
</div>
</td>
<td class="px-4 py-2 text-sm text-gray-700">
<% if variant[:selected_options] && !variant[:selected_options].empty? %>
<% variant[:selected_options].each do |option| %>
<span class="inline-block bg-gray-200 text-gray-800 text-xs round
<%= option[:name] %>: <%= option[:value] %>
</span>
<% end %>
<% else %>
<span class="text-gray-500">No options</span>
<% end %>
</td>
<td class="px-4 py-2 text-sm">
<div class="hs-dropdown relative inline-flex">
<button id="hs-dropdown-custom-icon-trigger" type="button" class="hs-dr
<svg class="flex-none size-4 text-gray-600 dark:text-neutral-500" xml
</button>
<div class="hs-dropdown-menu transition-[opacity,margin] duration hs-dr
<div class="p-1 space-y-0.5">
<a href="/products/<%= @product[:id] %>/variants/<%= variant[:id] %>
<a class="flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm
Purchases


```

