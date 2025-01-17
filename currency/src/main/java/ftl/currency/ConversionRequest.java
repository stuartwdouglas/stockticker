package ftl.currency;

import java.math.BigDecimal;

public record ConversionRequest(String to, String from, BigDecimal amount) {
}
