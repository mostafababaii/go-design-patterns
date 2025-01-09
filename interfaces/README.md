This directory contains an example implementation of a design pattern discussed in the book "100 Go Mistakes". The code demonstrates how to define an interface on the consumer side and implement it on the producer side, following best practices for clean and maintainable code.

## Overview

The example consists of three main components:
1. **Producer**: Implements the `DataProvider` which provides data.
2. **Consumer**: Defines the `DataProcessorInterface` and implements the `DataConsumer` which processes the data.
3. **Main**: Creates instances of `DataProvider` and `DataConsumer` and processes the data.

## Why This Pattern?

The book "100 Go Mistakes" emphasizes that an interface should live on the consumer side in most cases. However, in particular contexts (for example, when we know—not foresee—that an abstraction will be helpful for consumers), we may want to have it on the producer side. If we do, we should strive to keep it as minimal as possible, increasing its reusability potential and making it more easily composable.

In this example, the interface DataProcessorInterface is defined on the consumer side, and the DataProvider implements this interface on the producer side. This design ensures that the interface is minimal, increasing its reusability potential and making it more easily composable.