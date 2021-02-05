# CO2 Sensor

## What

Use Arduino Nano 33 IoT, a CO2 sensor and a OLED display for displaying the current CO2 concentration.
The current vaue should also be exposed via http.
Stretch goal is to connect this to a openHab smarthome and show graphs in Grafana and trigger alerts for high values.

## Why

I want to know for how long I need to open windows in the morning to get fresh air into my house, without losing too
much heat in winter. Bonus: in the current situation something like this might be used as an indicator for the need
for fresh air for reducing the Coronavirus concentration.

## How

### Parts

- Arduino Nano 33 IoT
- Sensirion SCD30 CO2, humidity and temperature sensor module
- Waveshare 0.95 inch RGB OLED

### Some hints

- https://tinygo.org/microcontrollers/arduino-nano33-iot/

- for IntelliJ:
    - see https://medium.com/@brandentimm/arduino-development-with-tinygo-in-jetbrains-ides-f29dbf5da675
    - and disable module support in the project!

## License

Copyright 2021 Marc Sluiter

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.