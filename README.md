# colorconv

`colorconv` is a simple command-line utility for converting between RGB, HEX, HSL, and HSV colour spaces.

## Features

- **Convert from RGB**: Translates values such as `rgb(255, 99, 71)` into HEX, HSL, and HSV.
- **Convert from HEX**: Translates values like `#ff6347` into RGB, HSL, and HSV.
- **Convert from HSL**: Translates values such as `hsl(9, 100%, 64%)` into RGB, HEX, and HSV.
- **Convert from HSV**: Translates values such as `hsv(211, 40.1%, 91.0%)` into RGB, HEX, and HSL.
- **Straightforward Interface**: Intuitive commands for ease of use.

## Installation

If you have a Go environment configured, you can install the tool with the following command:

```bash
go install github.com/WhiteRaven777/colorconv@latest
```

Alternatively, you may clone the repository and build it manually:

```bash
git clone https://github.com/WhiteRaven777/colorconv.git
cd colorconv
go build
```

## Usage

The basic command syntax is as follows:

```
Usage:
  colorconv rgb <red> <green> <blue>
  colorconv rgb <hex>
  colorconv hsl <hue> <saturation> <lightness>
  colorconv hsv <hue> <saturation> <value>

Note:
  - <red>, <green>, <blue>: integers between 0 and 255
  - <hex>: 6-digit hexadecimal code (e.g. '#ff00ff' or 'ff00ff')
  - <hue>: floating point value between 0.0 and 360.0 degrees
  - <saturation>, <lightness>, <value>: floating point values from 0.0 to 100.0 per cent
```

### Examples
#### 1. Convert from RGB

```bash
# Convert skyblue
$ colorconv rgb 139 184 232
```

**Output:**
```
RGB: 139, 184, 232 [ #8bb8e8 ]
HSL: 211.0, 0.669, 0.727 [ hsl(211.0, 66.9%, 72.7%) ]
HSV: 211.0, 0.401, 0.910 [ hsv(211.0, 40.1%, 91.0%) ]
```

#### 2. Convert from HEX

The leading `#` is optional.

```bash
$ colorconv rgb 8bb8e8
```

**Output:**
```
RGB: 139, 184, 232 [ #8bb8e8 ]
HSL: 211.0, 0.669, 0.727 [ hsl(211.0, 66.9%, 72.7%) ]
HSV: 211.0, 0.401, 0.910 [ hsv(211.0, 40.1%, 91.0%) ]
```

#### 3. Convert from HSL

```bash
$ colorconv hsl 211 66.9 72.7
```

**Output:**
```
RGB: 139, 184, 232 [ #8bb8e8 ]
HSL: 211.0, 0.669, 0.727 [ hsl(211.0, 66.9%, 72.7%) ]
HSV: 211.0, 0.401, 0.910 [ hsv(211.0, 40.1%, 91.0%) ]
```

#### 4. Convert from HSV

```bash
$ colorconv hsv 211 40.17 91.0
```

**Output:**
```
RGB: 139, 184, 232 [ #8bb8e8 ]
HSL: 211.0, 0.669, 0.727 [ hsl(211.0, 66.9%, 72.7%) ]
HSV: 211.0, 0.401, 0.910 [ hsv(211.0, 40.1%, 91.0%) ]
```