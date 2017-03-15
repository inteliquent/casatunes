package casatunes

type RESTMediaItem struct {
  Flags int
  ID string
  GroupName string
  Title string
  Biography string
  Type int
  ArtworkURI string
  ArtworkRatio float32
  Album string
  Artists string
  Duration float64
  Genres string
  Released string
  Track string
  BitRate int
  ListenerCount int
  StationBand string
  StationDescription string
  StationFrequency string
  StationName string
  Rating int
  Description string
  ServiceName string
  ServiceLogoURI string
  TotalItems int
  SearchPlaceholderText string
  DisplayInfo []string
  ContextMenuItems []RESTContextMenuItem
}

type RESTContextMenuItem struct {
  Type int
  Title string
  Value string
}

type RESTMediaCollectionItem struct {
  Flags int
  ID string
  Title string
  Description string
  Biography string
  DisplayInfo []string
  ArtworkURI string
  ArtworkRatio float32
  SearchPlaceholderText string
  StartIndex int
  TotalAvailable int
  MediaItems []RESTMediaItem
  ContextMenuItems []RESTContextMenuItem
}
