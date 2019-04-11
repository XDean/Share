# SWT Utils

## Pretty Pack Table

```java
public static void autoPrettyPack(Table table) {
  prettyPack(table);
  table.addListener(SWT.Resize, e -> prettyPack(table));
  if ((table.getStyle() & SWT.VIRTUAL) != 0) {
    table.addListener(SWT.SetData, e -> prettyPack(table));
  }
}

public static void prettyPack(Table table) {
  int width = table.getClientArea().width;
  for (TableColumn column : table.getColumns()) {
    column.pack();
  }
  Map<TableColumn, Integer> widths = new HashMap<>();
  int count = table.getColumnCount();
  int widthNeed = Arrays.stream(table.getColumns()).mapToInt(c -> {
    int w = c.getWidth();
    widths.put(c, w);
    return w;
  }).sum();
  int each = (width - widthNeed) / count;
  // `getClientArea` is magic, do not delete it
  table.getClientArea();
  if (each > 0) {
    Arrays.stream(table.getColumns()).forEach(c -> c.setWidth(widths.get(c) + each));
  }
  table.getClientArea();
}
```

## Rotate X label

```java
public static double rotateTopAlign(GC gc, double width, String[] labels) {
  if (labels.length == 0) {
    return 0;
  }
  gc = gc == null ? DEFAULT_GC : gc;
  double gap = width / labels.length;
  Point[] sizes = Arrays.stream(labels).map(gc::textExtent).toArray(Point[]::new);
  return IntStream.range(0, sizes.length + 1)
      .mapToDouble(i -> {
        Point left;
        Point right;
        if (i == 0) {
          left = (right = sizes[0]);
        } else if (i == sizes.length) {
          left = (right = sizes[sizes.length - 1]);
        } else {
          left = sizes[i - 1];
          right = sizes[i];
        }
        if (left.x + right.x <= 2 * gap) {
          return 0;
        }
        double w1 = left.x;
        double h1 = left.y;
        double w2 = right.x;
        double h2 = right.y;
        double hAngle = findRoot(a -> gap * sin(a) + ((h2 * cos(a) + w2 * sin(a) / 2) -
            (h1 * cos(a) + w1 * sin(a)) / 2) * cos(a) - (h1 + h2) / 2 - 5,
            0, Math.PI / 2, 0.1);
        return hAngle;
      }).max()
      .orElse(0d) / Math.PI * 180;
}

public static double rotateCenterAlign(GC gc, double width, String[] labels) {
  if (labels.length == 0) {
    return 0;
  }
  gc = gc == null ? DEFAULT_GC : gc;
  double gap = width / labels.length;
  Point[] sizes = Arrays.stream(labels).map(gc::textExtent).toArray(Point[]::new);
  return IntStream.range(0, sizes.length + 1)
      .mapToDouble(i -> {
        Point left;
        Point right;
        if (i == 0) {
          left = (right = sizes[0]);
        } else if (i == sizes.length) {
          left = (right = sizes[sizes.length - 1]);
        } else {
          left = sizes[i - 1];
          right = sizes[i];
        }
        if (left.x + right.x <= 2 * gap) {
          return 0;
        }
        double leftAngle = Math.atan(left.y / ((double) left.x));
        double rightAngle = Math.atan(right.y / ((double) right.x));

        double r1 = Math.acos((left.x + right.x) / (2 * gap));
        if (r1 <= leftAngle && r1 <= rightAngle) {
          return r1;
        }
        if (leftAngle > rightAngle) {
          double r2 = findRoot(d -> 2 * Math.sin(d) * gap - right.y - left.y * Math.tan(d), rightAngle, leftAngle, 0.1);
          if (r2 >= rightAngle && r2 <= leftAngle) {
            return r2;
          }
        } else {
          double r2 = findRoot(d -> 2 * Math.sin(d) * gap - left.y - right.y * Math.tan(d), rightAngle, leftAngle, 0.1);
          if (r2 >= leftAngle && r2 <= rightAngle) {
            return r2;
          }
        }
        double r3 = Math.asin((left.y + right.y) / (2 * gap));
        if (!Double.isNaN(r3)) {
          return r3;
        }
        return Math.PI / 2;
      }).max()
      .orElse(0d) / Math.PI * 180;
}

public static void alignLabelWidth(GC gc, String[] labels) {
  gc = gc == null ? DEFAULT_GC : gc;
  Point[] sizes = Arrays.stream(labels).map(gc::textExtent).toArray(Point[]::new);
  int space = gc.textExtent(" ").x;
  int maxWidth = Arrays.stream(sizes).mapToInt(p -> p.x).max().orElse(0);
  for (int i = 0; i < sizes.length; i++) {
    int toFill = (maxWidth - sizes[i].x) / space;
    int left = toFill / 2;
    int right = toFill - left;
    labels[i] = repeat(" ", left) + labels[i] + repeat(" ", right);
  }
}

public static String repeat(String st, int times) {
  StringBuilder sb = new StringBuilder();
  while (times-- > 0) {
    sb.append(st);
  }
  return sb.toString();
}

/**
 * Find monotonically increasing function approximate root(right side), suppose y of from is
 * negative and y of to is positive
 */
public static double findRoot(DoubleUnaryOperator function, double from, double to, double precision) {
  if (function.applyAsDouble(from) > 0 || function.applyAsDouble(to) < 0) {
    return Double.NaN;
  }
  double f = from;
  double t = to;
  double m = (f + t) / 2;
  while (true) {
    double v = function.applyAsDouble(m);
    if ((v >= 0 && v < precision) || (t - f) < 1e-6) {
      return m;
    }
    if (v > 0) {
      t = m;
    } else {
      f = m;
    }
    m = (f + t) / 2;
  }
}
```