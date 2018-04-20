<?php
include_once('geoPHP.inc');
$zipcodes = json_decode(file_get_contents('https://dawa.aws.dk/postnumre'));
$out = [];
foreach ($zipcodes as $zipcode) {
    $geojson = file_get_contents($zipcode->href . '?format=geojson');
    try {
        $polygon = geoPHP::load($geojson, 'json');
        $centroid = $polygon->getCentroid();
        $out[] = [
        'zip' => $zipcode->nr,
        'lat' => $centroid->getX(),
        'lon' => $centroid->getY(),
        ];
    } catch (Exception $e) {
    }
}
echo json_encode($out);
