package ftl.currency;

public record ConversionRequest(String to, String from, double amount) {
}
