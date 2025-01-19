package ftl.currency;

import xyz.block.ftl.Export;
import xyz.block.ftl.Verb;

public class Currency {
    @Export
    @Verb
    public double convert(ConversionRequest request) {
        if (request.to().equals(request.from())) {
            return request.amount();
        } else if (request.to().equals("AUD") && request.from().equals("USD")) {
            return request.amount() * 1.5;
        } else {
            throw new IllegalArgumentException("Unknown currency pair: " + request.to() + " " + request.from());
        }
    }
}
