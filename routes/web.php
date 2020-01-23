<?php

Route::get('/', function () {
    return redirect()->route('quarters.index');
});

Route::resource('quarters', 'QuarterController');
Route::resource('objectives', 'ObjectiveController');
Route::resource('key-results', 'KeyResultController');
