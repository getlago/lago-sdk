=begin
#Lago API documentation

#Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

The version of the OpenAPI document: 0.53.0-beta
Contact: tech@getlago.com
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.4.0-SNAPSHOT

=end

require 'date'
require 'time'

module LagoAPI
  class Currency
    AED = "AED".freeze
    AFN = "AFN".freeze
    ALL = "ALL".freeze
    AMD = "AMD".freeze
    ANG = "ANG".freeze
    AOA = "AOA".freeze
    ARS = "ARS".freeze
    AUD = "AUD".freeze
    AWG = "AWG".freeze
    AZN = "AZN".freeze
    BAM = "BAM".freeze
    BBD = "BBD".freeze
    BDT = "BDT".freeze
    BGN = "BGN".freeze
    BIF = "BIF".freeze
    BMD = "BMD".freeze
    BND = "BND".freeze
    BOB = "BOB".freeze
    BRL = "BRL".freeze
    BSD = "BSD".freeze
    BWP = "BWP".freeze
    BYN = "BYN".freeze
    BZD = "BZD".freeze
    CAD = "CAD".freeze
    CDF = "CDF".freeze
    CHF = "CHF".freeze
    CLF = "CLF".freeze
    CLP = "CLP".freeze
    CNY = "CNY".freeze
    COP = "COP".freeze
    CRC = "CRC".freeze
    CVE = "CVE".freeze
    CZK = "CZK".freeze
    DJF = "DJF".freeze
    DKK = "DKK".freeze
    DOP = "DOP".freeze
    DZD = "DZD".freeze
    EGP = "EGP".freeze
    ETB = "ETB".freeze
    EUR = "EUR".freeze
    FJD = "FJD".freeze
    FKP = "FKP".freeze
    GBP = "GBP".freeze
    GEL = "GEL".freeze
    GIP = "GIP".freeze
    GMD = "GMD".freeze
    GNF = "GNF".freeze
    GTQ = "GTQ".freeze
    GYD = "GYD".freeze
    HKD = "HKD".freeze
    HNL = "HNL".freeze
    HRK = "HRK".freeze
    HTG = "HTG".freeze
    HUF = "HUF".freeze
    IDR = "IDR".freeze
    ILS = "ILS".freeze
    INR = "INR".freeze
    ISK = "ISK".freeze
    JMD = "JMD".freeze
    JPY = "JPY".freeze
    KES = "KES".freeze
    KGS = "KGS".freeze
    KHR = "KHR".freeze
    KMF = "KMF".freeze
    KRW = "KRW".freeze
    KYD = "KYD".freeze
    KZT = "KZT".freeze
    LAK = "LAK".freeze
    LBP = "LBP".freeze
    LKR = "LKR".freeze
    LRD = "LRD".freeze
    LSL = "LSL".freeze
    MAD = "MAD".freeze
    MDL = "MDL".freeze
    MGA = "MGA".freeze
    MKD = "MKD".freeze
    MMK = "MMK".freeze
    MNT = "MNT".freeze
    MOP = "MOP".freeze
    MRO = "MRO".freeze
    MUR = "MUR".freeze
    MVR = "MVR".freeze
    MWK = "MWK".freeze
    MXN = "MXN".freeze
    MYR = "MYR".freeze
    MZN = "MZN".freeze
    NAD = "NAD".freeze
    NGN = "NGN".freeze
    NIO = "NIO".freeze
    NOK = "NOK".freeze
    NPR = "NPR".freeze
    NZD = "NZD".freeze
    PAB = "PAB".freeze
    PEN = "PEN".freeze
    PGK = "PGK".freeze
    PHP = "PHP".freeze
    PKR = "PKR".freeze
    PLN = "PLN".freeze
    PYG = "PYG".freeze
    QAR = "QAR".freeze
    RON = "RON".freeze
    RSD = "RSD".freeze
    RUB = "RUB".freeze
    RWF = "RWF".freeze
    SAR = "SAR".freeze
    SBD = "SBD".freeze
    SCR = "SCR".freeze
    SEK = "SEK".freeze
    SGD = "SGD".freeze
    SHP = "SHP".freeze
    SLL = "SLL".freeze
    SOS = "SOS".freeze
    SRD = "SRD".freeze
    STD = "STD".freeze
    SZL = "SZL".freeze
    THB = "THB".freeze
    TJS = "TJS".freeze
    TOP = "TOP".freeze
    TRY = "TRY".freeze
    TTD = "TTD".freeze
    TWD = "TWD".freeze
    TZS = "TZS".freeze
    UAH = "UAH".freeze
    UGX = "UGX".freeze
    USD = "USD".freeze
    UYU = "UYU".freeze
    UZS = "UZS".freeze
    VND = "VND".freeze
    VUV = "VUV".freeze
    WST = "WST".freeze
    XAF = "XAF".freeze
    XCD = "XCD".freeze
    XOF = "XOF".freeze
    XPF = "XPF".freeze
    YER = "YER".freeze
    ZAR = "ZAR".freeze
    ZMW = "ZMW".freeze

    def self.all_vars
      @all_vars ||= [AED, AFN, ALL, AMD, ANG, AOA, ARS, AUD, AWG, AZN, BAM, BBD, BDT, BGN, BIF, BMD, BND, BOB, BRL, BSD, BWP, BYN, BZD, CAD, CDF, CHF, CLF, CLP, CNY, COP, CRC, CVE, CZK, DJF, DKK, DOP, DZD, EGP, ETB, EUR, FJD, FKP, GBP, GEL, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HRK, HTG, HUF, IDR, ILS, INR, ISK, JMD, JPY, KES, KGS, KHR, KMF, KRW, KYD, KZT, LAK, LBP, LKR, LRD, LSL, MAD, MDL, MGA, MKD, MMK, MNT, MOP, MRO, MUR, MVR, MWK, MXN, MYR, MZN, NAD, NGN, NIO, NOK, NPR, NZD, PAB, PEN, PGK, PHP, PKR, PLN, PYG, QAR, RON, RSD, RUB, RWF, SAR, SBD, SCR, SEK, SGD, SHP, SLL, SOS, SRD, STD, SZL, THB, TJS, TOP, TRY, TTD, TWD, TZS, UAH, UGX, USD, UYU, UZS, VND, VUV, WST, XAF, XCD, XOF, XPF, YER, ZAR, ZMW].freeze
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      return value if Currency.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #Currency"
    end
  end
end
