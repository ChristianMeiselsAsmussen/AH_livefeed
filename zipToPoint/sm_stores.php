<?php
$stores = json_decode(file_get_contents('https://sportmaster.dk/api/stores/all'));
$out = [];
foreach ($stores as $store) {
  if ($store->type === 'store') {
    $out[] = [
      'Id' => $store->id,
      'Lat' => $store->location->lat,
      'Lon' => $store->location->lon,
    ];
  }
}
echo json_encode($out);
